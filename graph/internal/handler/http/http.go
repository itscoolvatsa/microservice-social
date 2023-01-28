package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"microservice/graph/internal/controller"
	"microservice/pkg/errors"
	"microservice/pkg/middleware"
	errresp "microservice/pkg/response/error"
	"net/http"
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
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000/, http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "PUT", "HEAD"},
		AllowHeaders:     []string{"Content-Type, Access-Control-Allow-Headers, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type, Access-Control-Allow-Headers, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000/" || origin == "http://localhost:3000"
		},
	})

	incomingRoutes.NoRoute(h.NotFound())
	incomingRoutes.Use(c)
	routes := incomingRoutes.Group("/api")
	{
		routes.POST("/graph/add-follow/:following", middleware.CurrentUser(symmetricKey), h.AddFollow())
	}
}

func (h *HttpHandler) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		errresp.New(http.StatusNotFound, false, errors.ErrNotFound).SendResponse(c)
	}
}
