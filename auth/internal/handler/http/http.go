package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"microservice/auth/internal/controller"
	"microservice/pkg/middleware"
	jsonresp "microservice/pkg/response/json"
	"microservice/pkg/token"
	"net/http"
)

// Handler for user service http gateway.
type Handler struct {
	ctrl       *controller.Controller
	tokenMaker token.Maker
}

// New creates a new user http handler.
func New(ctrl *controller.Controller, symmetricKey string) (*Handler, error) {
	maker, err := token.New(symmetricKey)
	if err != nil {
		return nil, err
	}
	return &Handler{
		ctrl:       ctrl,
		tokenMaker: maker,
	}, nil
}

var validate = validator.New()

// SignOutUser Logs out the user
func (h *Handler) SignOutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Authorization", "")
		resp := jsonresp.New(http.StatusAccepted, true, "", "logged out successfully")
		resp.SendResponse(c)
	}
}

func (h *Handler) UserRoutes(incomingRoutes *gin.Engine, symmetricKey string, NATS_URL string) {
	incomingRoutes.POST("/users/signup", h.SignupUser(NATS_URL))
	incomingRoutes.POST("/users/signin", h.SignInUser())
	incomingRoutes.GET("/users/signout", h.SignOutUser())
	incomingRoutes.GET("/users/current-user", middleware.CurrentUser(symmetricKey), h.CurrentUser())
}
