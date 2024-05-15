package main

import (
	"hillel_auction/config"
	"hillel_auction/db"
	"hillel_auction/logger"
	"hillel_auction/repository"
	"hillel_auction/server"
	"hillel_auction/server/handlers"
	"hillel_auction/services"
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

	database := db.New(l, cfg)
	err = database.Connect()
	if err != nil {
		l.Fatal("failed to connect to the database", err)
	}

	userRepos := repository.NewUsersRepository(database, l)

	userService := services.NewUserService(l, userRepos)

	h := handlers.NewHandlers(l, userService)

	s := server.NewServer(cfg, h)
	err = s.Start()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
