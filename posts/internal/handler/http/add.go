package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"microservice/pkg/token"
	"microservice/posts/pkg/model"
	"net/http"
	"time"
)

// AddPost handles post request for signup
func (h *Handler) AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		payload := c.MustGet("access").(*token.Payload)

		var post model.Post

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&post); err != nil {
			errresp.New(http.StatusBadRequest, false, err).
				SendResponse(c)
			return
		}

		//validate the data based on user struct
		validationErr := validate.Struct(post)

		if validationErr != nil {
			errresp.New(http.StatusBadRequest, false, validationErr).
				SendResponse(c)
			return
		}

		defer cancel()

		post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UserId = payload.UserId

		//	saving the suer to the database
		err := h.ctrl.AddPost(ctx, &post)

		if err != nil {
			errresp.New(http.StatusBadRequest, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		jsonresp.New(http.StatusAccepted, true, post, "post added successfully").
			SendResponse(c)
	}
}
