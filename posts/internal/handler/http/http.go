package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"microservice/pkg/errors"
	"microservice/pkg/middleware"
	errresp "microservice/pkg/response/error"
	"microservice/posts/internal/controller"
	"net/http"
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
		routes.POST("/posts/add", middleware.CurrentUser(symmetricKey), h.AddPost())
		routes.GET("/posts/get-posts", middleware.CurrentUser(symmetricKey), h.GetPostsForUser())
		routes.GET("/posts/get-media/:fid", h.GetMedia())
	}
}

func (h *Handler) NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		errresp.New(http.StatusNotFound, false, errors.ErrNotFound).SendResponse(c)
	}
}
