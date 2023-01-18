package http

import (
	"github.com/gin-gonic/gin"
	"microservice/graph/internal/controller"
	"microservice/pkg/middleware"
)

// HttpHandler for user service http gateway.
type HttpHandler struct {
	ctrl *controller.Controller
}

// New creates a new user http handler.
func New(ctrl *controller.Controller) (*HttpHandler, error) {
	return &HttpHandler{
		ctrl: ctrl,
	}, nil
}

func (h *HttpHandler) GraphRoutes(incomingRoutes *gin.Engine, symmetricKey string) {
	incomingRoutes.POST("/graph/add-follow/:following", middleware.CurrentUser(symmetricKey), h.AddFollow())
}
