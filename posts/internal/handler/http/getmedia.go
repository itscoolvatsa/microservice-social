package http

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	"net/http"
)

// GetMedia returns media by the id provided
func (h *Handler) GetMedia() gin.HandlerFunc {
	return func(c *gin.Context) {
		prms := c.Param("fid")
		fid, err := primitive.ObjectIDFromHex(prms)

		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		file, err := h.ctrl.GetMedia(fid)

		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		c.Data(http.StatusOK, "image/jpeg", file.Bytes())
	}
}
