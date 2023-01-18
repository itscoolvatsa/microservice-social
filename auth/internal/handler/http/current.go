package http

import (
	"github.com/gin-gonic/gin"
	jsonresp "microservice/pkg/response/json"
	"microservice/pkg/token"
	"net/http"
)

func (h *Handler) CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := c.MustGet("access").(*token.Payload)
		resp := jsonresp.New(http.StatusOK, true, payload, "user logged in")
		resp.SendResponse(c)
	}
}
