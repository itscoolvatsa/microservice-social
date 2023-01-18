package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"microservice/pkg/token"
	"microservice/posts/pkg/model"
	"net/http"
	"time"
)

// AddPost handles post request for signup
func (h *Handler) AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		payload := c.MustGet("access").(*token.Payload)
		fmt.Printf("%v", payload)

		var post model.Post

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//validate the data based on user struct
		validationErr := validate.Struct(post)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		defer cancel()

		post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UserId = payload.UserId

		//	saving the suer to the database
		err := h.ctrl.AddPost(ctx, &post)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while adding the post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "post added successfully"})
	}
}
