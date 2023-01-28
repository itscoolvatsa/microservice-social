package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice/auth/internal/controller"
	httpHandler "microservice/auth/internal/handler/http"
	"microservice/auth/internal/repository/mongodb"
	config2 "microservice/auth/pkg/config"
)

func main() {
	config, _ := config2.LoadConfig(".")
	println(config.MONGOURI)

	repo := mongodb.New(config.MONGOURI)
	ctrl := controller.New(repo)

	httpServer(ctrl, config)
}

func httpServer(ctrl *controller.Controller, config config2.Config) {
	handler, err := httpHandler.New(ctrl, config.TokenSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}

	server := gin.New()
	handler.UserRoutes(server, config.TokenSymmetricKey, config.NATSURL)
	port := config.HttpServerAddress

	server.Run(port)
}
