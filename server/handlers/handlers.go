package handlers

import "hillel_auction/logger"

type Handlers struct {
	log *logger.Logger
}

func NewHandlers(log *logger.Logger) *Handlers {
	return &Handlers{
		log: log,
	}
}
