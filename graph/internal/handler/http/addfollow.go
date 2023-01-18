package http

import (
	"github.com/gin-gonic/gin"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
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
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		jsonresp.New(http.StatusAccepted, true, "", "relationship established").
			SendResponse(c)
	}
}
