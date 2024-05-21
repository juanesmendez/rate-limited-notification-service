package util

import (
	"NotificationService/enum"
	"strings"
)

func MapToNotificationType(notificationType string) enum.NotificationType {
	switch strings.ToLower(notificationType) {
	case "status":
		return enum.Status
	case "news":
		return enum.News
	case "marketing":
		return enum.Marketing
	default:
		return enum.Unknown
	}
}
