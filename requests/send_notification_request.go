package requests

type SendNotificationRequest struct {
	UserId           string `json:"user_id" validate:"required"`
	NotificationType string `json:"notification_type" validate:"required"`
	Message          string `json:"message" validate:"required"`
}
