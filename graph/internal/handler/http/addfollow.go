package http

import (
	"github.com/gin-gonic/gin"
	"microservice/pkg/token"
	"net/http"
)

func (h *HttpHandler) AddFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := c.MustGet("access").(*token.Payload)
		following_id := c.Param("following")

		err := h.ctrl.AddRelationShip(c, payload.UserId.Hex(), following_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "relationship established"})
	}
}
