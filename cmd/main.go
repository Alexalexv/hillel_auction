package main

import (
	"hillel_auction/config"
	"hillel_auction/db"
	"hillel_auction/logger"
	"hillel_auction/server"
	"hillel_auction/server/handlers"
	"log"

	"hillel_auction/docs"
)

// @title           AuthServiceAPI
// @version         1.0
// @description     My first project.
// @termsOfService  http://swagger.io/terms/

// @BasePath /
func main() {
	docs.SwaggerInfo.Host = ""

	cfg, err := config.NewConfiguration()
	if err != nil {
		log.Fatal("failed to create configuration")
	}

	l := logger.NewLogger(cfg)
	h := handlers.NewHandlers(l)

	database := db.New(l, cfg)
	err = database.Connect()
	if err != nil {
		l.Fatal("failed to connect to the database", err)
	}

	s := server.NewServer(cfg, h)
	err = s.Start()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
