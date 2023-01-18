package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microservice/pkg/token"
	"net/http"
)

func (h *Handler) CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := c.MustGet("access").(*token.Payload)
		fmt.Printf("%v", payload)
		c.JSON(http.StatusOK, gin.H{"user": payload})
	}
}
