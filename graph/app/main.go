package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice/graph/internal/controller"
	"microservice/graph/internal/handler/http"
	natshandler "microservice/graph/internal/handler/natsmsg"
	"microservice/graph/internal/repository/natsmsg"
	"microservice/graph/internal/repository/neo4j"
	config2 "microservice/graph/pkg/config"
)

func main() {
	config, _ := config2.LoadConfig("./")
	println(config.NATSURL)
	println(config.NEO4JURI)
	userRepo := neo4j.New(config.NEO4JURI)
	natsMsg := natsmsg.New(config.NATSURL)
	ctrl := controller.New(userRepo, natsMsg)
	httpServer(ctrl, config)
}

func httpServer(ctrl *controller.Controller, config config2.Config) {
	handler, err := natshandler.New(ctrl)
	httphandler, err := http.New(ctrl)
	if err != nil {
		log.Fatal(err)
	}

	server := gin.New()
	httphandler.GraphRoutes(server, config.TokenSymmetricKey)
	handler.ReceivesMessage("user")

	port := config.HttpServerAddress

	server.Run(port)
}
