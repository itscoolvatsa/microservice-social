package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"microservice/auth/internal/repository/password"
	"microservice/auth/pkg/model"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
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
			resp := errresp.New(http.StatusBadRequest, false, err.Error())
			resp.SendResponse(c)
			return
		}

		//you'll check if the email has already been used by another user
		storedUser, err := h.ctrl.FindUser(ctx, providedUser.Email)

		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, "email or password is wrong")
			resp.SendResponse(c)
			return
		}
		defer cancel()

		check := password.ComparePassword(storedUser.Password, providedUser.Password)

		if !check {
			resp := errresp.New(http.StatusBadRequest, false, "email or password is wrong")
			resp.SendResponse(c)
			return
		}

		accessToken, _, err := h.tokenMaker.CreateToken(storedUser.Email, storedUser.ID.Hex(), time.Hour)
		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, "something went wrong")
			resp.SendResponse(c)
			return
		}

		accessToken = "Bearer " + accessToken
		c.Header("Authorization", accessToken)
		//c.JSON(http.StatusOK, gin.H{"message": storedUser})

		data := storedUser
		resp := jsonresp.New(http.StatusOK, true, data, "signed in successfully")
		resp.SendResponse(c)
	}
}
