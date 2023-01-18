package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservice/auth/internal/repository/natsmsg"
	"microservice/auth/internal/repository/password"
	"microservice/auth/pkg/model"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"net/http"
	"time"
)

// SignupUser handles post request for signup
func (h *Handler) SignupUser(NATS_URL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var user model.User

		//convert the JSON data coming from postman to something that golang understands
		if err := c.BindJSON(&user); err != nil {
			resp := errresp.New(http.StatusBadRequest, false, err)
			resp.SendResponse(c)
			return
		}

		//validate the data based on user struct
		validationErr := validate.Struct(user)

		if validationErr != nil {
			resp := errresp.New(http.StatusBadRequest, false, validationErr)
			resp.SendResponse(c)
			return
		}

		//you'll check if the email has already been used by another user
		count, err := h.ctrl.CountUser(ctx, user.Email)

		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrInternalServer)
			resp.SendResponse(c)
			return
		}

		if count > 0 {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrAlreadyExists)
			resp.SendResponse(c)
			return
		}

		// hashing the password before saving into the database
		hashedPassword, _ := password.HashPassword(user.Password)

		user.Password = hashedPassword
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		//	saving the suer to the database
		userId, err := h.ctrl.AddUser(ctx, &user)

		if err != nil {
			resp := errresp.New(http.StatusBadRequest, false, errors.ErrInternalServer)
			resp.SendResponse(c)
			return
		}

		println(NATS_URL)
		println(userId.InsertedID)
		natSrv := natsmsg.New(NATS_URL)
		natSrv.SendMessage(natsmsg.NatsUser{
			UserId: userId.InsertedID.(primitive.ObjectID),
			Name:   user.Name,
		})

		resp := jsonresp.New(http.StatusBadRequest, true, "", "user signed up successfully")
		resp.SendResponse(c)
	}
}
