package handlers

import (
	"hillel_auction/logger"
	"hillel_auction/services"
)

type Handlers struct {
	log         *logger.Logger
	userService *services.UserService
}

func NewHandlers(log *logger.Logger, userService *services.UserService) *Handlers {
	return &Handlers{
		log:         log,
		userService: userService,
	}
}
