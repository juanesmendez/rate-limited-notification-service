package redis

import (
	"NotificationService/environment"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", environment.RedisHost, environment.RedisPort),
		Password: environment.RedisPassword,
	})
}

func TestRedisConnection(client *redis.Client) {
	pong, err := client.Ping().Result()

	if err != nil {
		log.Println("Error connecting to Redis:", err)
		return
	}
	log.Println("Connected to Redis:", pong)
}
