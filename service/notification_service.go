package service

import (
	"NotificationService/enum"
	"NotificationService/util"
	"errors"
	"fmt"
	"log"
	"strings"
)

type NotificationService interface {
	Send(notificationType string, userId string, message string) error
}

type notificationServiceImpl struct {
	gateway     GatewayService
	rateLimiter RateLimiterService
}

func NewNotificationServiceImpl(gateway GatewayService, rateLimiter RateLimiterService) NotificationService {
	return &notificationServiceImpl{
		gateway:     gateway,
		rateLimiter: rateLimiter,
	}
}

func (service *notificationServiceImpl) Send(notificationTypeStr string, userId string, message string) error {
	notificationType := util.MapToNotificationType(notificationTypeStr)

	if notificationType == enum.Unknown {
		log.Println("[NotificationService] unknown notification_type...")

		return errors.New("unknown notification type")
	}

	canSend, err := service.rateLimiter.CanSend(userId, notificationType)
	if err != nil {

		return fmt.Errorf("error checking if message can be sent: %w", err)
	}

	if !canSend {
		log.Printf("[NotificationService] rate limit exceeded for user=%s and notificationType=%s", userId, notificationType)

		return fmt.Errorf("rate limit exceeded for user=%s and notificationType=%s", userId, notificationType)
	}

	formattedMessage := fmt.Sprintf("[%s] %s\n", strings.ToUpper(notificationTypeStr), message)
	service.gateway.Send(userId, formattedMessage)

	return nil
}
