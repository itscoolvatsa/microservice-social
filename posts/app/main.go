package main

import (
	"github.com/gin-gonic/gin"
	"log"
	config2 "microservice/auth/pkg/config"
	"microservice/posts/internal/controller"
	httpHandler "microservice/posts/internal/handler/http"
	"microservice/posts/internal/repository/mongodb"
)

func main() {
	config, _ := config2.LoadConfig("./")
	println("mongo")
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
	handler.PostRoutes(server, config.TokenSymmetricKey)

	port := config.HttpServerAddress

	server.Run(port)
}
