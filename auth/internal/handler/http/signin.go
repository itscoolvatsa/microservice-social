package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"microservice/auth/internal/repository/password"
	"microservice/auth/pkg/model"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"net/http"
	"time"
)

// SignInUser handles post request for signup
func (h *Handler) SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		var storedUser *model.User
		var providedUser model.User

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&providedUser); err != nil {
			resp := errresp.New(http.StatusBadRequest, false, err)
			resp.SendResponse(c)
			return
		}

		//you'll check if the email has already been used by another user
		storedUser, err := h.ctrl.FindUser(ctx, providedUser.Email)

		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrInvalidCredentials)
			resp.SendResponse(c)
			return
		}

		check := password.ComparePassword(storedUser.Password, providedUser.Password)

		if !check {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrInvalidCredentials)
			resp.SendResponse(c)
			return
		}

		accessToken, _, err := h.tokenMaker.CreateToken(storedUser.Email, storedUser.ID.Hex(), time.Hour)
		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrInternalServer)
			resp.SendResponse(c)
			return
		}

		accessToken = "Bearer " + accessToken
		c.Header("Authorization", accessToken)

		data := storedUser
		resp := jsonresp.New(http.StatusOK, true, data, "signed in successfully")
		resp.SendResponse(c)
	}
}
