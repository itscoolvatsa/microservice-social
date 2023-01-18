package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"microservice/auth/internal/repository/password"
	"microservice/auth/pkg/model"
	"net/http"
	"time"
)

// SignInUser handles post request for signup
func (h *Handler) SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

		var storedUser *model.User
		var providedUser model.User

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&providedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//you'll check if the email has already been used by another user
		storedUser, err := h.ctrl.FindUser(ctx, providedUser.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}
		defer cancel()

		check := password.ComparePassword(storedUser.Password, providedUser.Password)

		if !check {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is wrong"})
			return
		}

		accessToken, _, err := h.tokenMaker.CreateToken(storedUser.Email, storedUser.ID.Hex(), time.Hour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while token generation"})
			return
		}

		accessToken = "Bearer " + accessToken
		c.Header("Authorization", accessToken)
		c.JSON(http.StatusOK, gin.H{"message": storedUser})
	}
}
