package constant

import (
	"NotificationService/enum"
	"NotificationService/model"
	"time"
)

var NotificationTypeRates = map[enum.NotificationType]model.RateLimitRule{
	enum.Status: {
		Limit:      2,
		TimeWindow: time.Minute * 2,
	},
	enum.News: {
		Limit:      1,
		TimeWindow: time.Minute * 60 * 24,
	},
	enum.Marketing: {
		Limit:      3,
		TimeWindow: time.Minute * 60 * 3,
	},
}
