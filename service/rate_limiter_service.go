package service

import (
	"NotificationService/constant"
	"NotificationService/enum"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type RateLimiterService interface {
	CanSend(userId string, notificationType enum.NotificationType) (bool, error)
}

type RateLimiterServiceImpl struct {
	redisClient *redis.Client
}

func NewRateLimiterImpl(redisClient *redis.Client) RateLimiterService {
	return &RateLimiterServiceImpl{
		redisClient: redisClient,
	}
}

func (service *RateLimiterServiceImpl) CanSend(userId string, notificationType enum.NotificationType) (bool, error) {
	log.Printf("[RateLimiterService] notificationType: %s\n", notificationType)
	key := fmt.Sprintf("%s:%s:%s", constant.RedisKeyPrefix, userId, notificationType)
	numMessagesStr, err := service.redisClient.Get(key).Result()

	if err != nil || numMessagesStr == "" {

		return service.handleKeyNotFound(userId, notificationType, key)
	}
	numMessages, err := strconv.Atoi(numMessagesStr)

	if err != nil {
		log.Printf("[RateLimiterService] error parsing redis key value to integer\n")

		return false, errors.New("error parsing redis key value to integer")
	}

	if numMessages < constant.NotificationTypeRates[notificationType].Limit {

		return service.handleValidNumMessages(userId, notificationType, key, numMessages)
	}

	return false, nil
}

func (service *RateLimiterServiceImpl) handleKeyNotFound(userId string, notificationType enum.NotificationType, key string) (bool, error) {
	log.Printf("[RateLimiterService] 0 messages have been sent to user %s of type %s\n", userId, notificationType)
	value := 1
	err := service.redisClient.Set(key, value, constant.NotificationTypeRates[notificationType].TimeWindow).Err()

	if err != nil {
		log.Println("[RateLimiterService] error setting redis key-value pair")
		return false, errors.New("error setting redis key-value pair")
	}
	log.Printf("[RateLimiterService] set redis key-value pair: key=%s value=%d\n", key, value)

	return true, nil
}

func (service *RateLimiterServiceImpl) handleValidNumMessages(
	userId string,
	notificationType enum.NotificationType,
	key string,
	numMessages int,
) (bool, error) {
	log.Printf("[RateLimiterService] %d messages have been sent for user %s of type %s\n", numMessages, userId, notificationType)
	ttl, err := service.redisClient.TTL(key).Result()

	if err != nil {
		log.Printf("error getting existing key ttl")

		return false, errors.New("error getting existing key ttl")
	}
	numMessages++
	err = service.redisClient.Set(key, numMessages, ttl).Err()

	if err != nil {
		log.Println("[RateLimiterService] error setting redis key-value pair")

		return false, errors.New("error setting redis key-value pair")
	}
	log.Printf("[RateLimiterService] can send message of type %s for user %s, for a new total of %d\n", notificationType, userId, numMessages)

	return true, nil
}
