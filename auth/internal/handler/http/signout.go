package http

import (
	"github.com/gin-gonic/gin"
	jsonresp "microservice/pkg/response/json"
	"net/http"
)

// SignOutUser Logs out the user
func (h *Handler) SignOutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Authorization", "")
		resp := jsonresp.New(http.StatusAccepted, true, "", "logged out successfully")
		resp.SendResponse(c)
	}
}
