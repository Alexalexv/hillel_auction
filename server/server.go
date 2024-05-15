package server

import (
	"hillel_auction/config"
	"hillel_auction/server/handlers"
	"hillel_auction/server/midlewares"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	e    *echo.Echo
	port string
}

func NewServer(cfg *config.Configuration, handlers *handlers.Handlers) *Server {
	e := echo.New()

	e.POST("/signup", handlers.SighUp)
	e.POST("/signin", handlers.SignIn)
	e.POST("/refresh", handlers.RefreshJWT)

	authGroup := e.Group("", midlewares.Authorization)

	authGroup.POST("/items", handlers.CreateItem)
	authGroup.GET("/items/all", handlers.GetAllItems)
	authGroup.GET("/items/:id", handlers.GetItem)
	authGroup.DELETE("/items/:id", handlers.DeleteItem)
	authGroup.PUT("/items/:id", handlers.UpdateItem)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &Server{
		e:    e,
		port: cfg.Port,
	}
}

func (s Server) Start() error {
	return s.e.Start(s.port)
}
