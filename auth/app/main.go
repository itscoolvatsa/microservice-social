package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice/auth/internal/controller"
	httpHandler "microservice/auth/internal/handler/http"
	"microservice/auth/internal/repository/mongodb"
	config2 "microservice/auth/pkg/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
	server.Use(CORSMiddleware())
	port := config.HttpServerAddress

	server.Run(port)
}
