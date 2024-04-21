package server

import (
	"hillel_auction/config"
	"hillel_auction/server/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	e    *echo.Echo
	port string
}

func NewServer(cfg *config.Configuration, handlers *handlers.Handlers) *Server {
	e := echo.New()

	e.POST("/item", handlers.CreateItem)
	e.GET("/item/all", handlers.GetAllItems)
	e.GET("/item/:id", handlers.GetItem)
	e.DELETE("/item/:id", handlers.DeleteItem)
	e.PUT("/item/:id", handlers.UpdateItem)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &Server{
		e:    e,
		port: cfg.Port,
	}
}

func (s Server) Start() error {
	return s.e.Start(s.port)
}
