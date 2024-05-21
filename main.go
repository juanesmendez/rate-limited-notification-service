package main

import (
	"NotificationService/controller"
	"NotificationService/redis"
	"NotificationService/router"
	"NotificationService/service"
)

func main() {
	redisClient := redis.NewRedisClient()
	redis.TestRedisConnection(redisClient)

	gatewayService := service.NewGatewayServiceImpl()
	rateLimiterService := service.NewRateLimiterImpl(redisClient)
	notificationService := service.NewNotificationServiceImpl(gatewayService, rateLimiterService)

	notificationController := controller.NewNotificationControllerImpl(notificationService)

	serverRouter := router.NewRouter(notificationController)
	serverRouter.Serve()
}
