package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"microservice/pkg/middleware"
	"microservice/posts/internal/controller"
)

// Handler for user service http gateway.
type Handler struct {
	ctrl *controller.Controller
}

// New creates a new user http handler.
func New(ctrl *controller.Controller, symmetricKey string) (*Handler, error) {
	return &Handler{
		ctrl: ctrl,
	}, nil
}

var validate = validator.New()

func (h *Handler) PostRoutes(incomingRoutes *gin.Engine, symmetricKey string) {
	incomingRoutes.POST("/posts/add", middleware.CurrentUser(symmetricKey), h.AddPost())
}
