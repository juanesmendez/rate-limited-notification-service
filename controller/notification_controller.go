package controller

import (
	"NotificationService/requests"
	"NotificationService/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NotificationController interface {
	Send(c echo.Context) error
}

type notificationControllerImpl struct {
	notificationService service.NotificationService
}

func NewNotificationControllerImpl(notificationService service.NotificationService) NotificationController {
	return &notificationControllerImpl{
		notificationService: notificationService,
	}
}

func (ctr *notificationControllerImpl) Send(c echo.Context) error {
	var sendNotificationBody requests.SendNotificationRequest

	if err := c.Bind(&sendNotificationBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request body",
		})
	}

	if err := validator.New().Struct(sendNotificationBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	err := ctr.notificationService.Send(sendNotificationBody.NotificationType, sendNotificationBody.UserId, sendNotificationBody.Message)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "email sent",
	})
}
