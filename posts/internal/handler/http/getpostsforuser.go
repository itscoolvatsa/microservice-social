package http

import (
	"github.com/gin-gonic/gin"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"microservice/pkg/token"
	"net/http"
)

// GetPostsForUser returns all posts for the user in descending order
func (h *Handler) GetPostsForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := c.MustGet("access").(*token.Payload)

		posts, err := h.ctrl.GetPostsForUser(payload.UserId)

		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		jsonresp.New(http.StatusOK, true, posts, "posts send successfully").SendResponse(c)
	}
}
