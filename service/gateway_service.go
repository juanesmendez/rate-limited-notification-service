package service

import "log"

type GatewayService interface {
	Send(userId string, message string)
}

type gatewayServiceImpl struct {
}

func NewGatewayServiceImpl() GatewayService {
	return &gatewayServiceImpl{}
}

func (service *gatewayServiceImpl) Send(userId string, message string) {
	log.Printf("[GatewayService] sending message to user %s: %s\n", userId, message)
}
