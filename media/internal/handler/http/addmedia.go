package http

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"net/http"
	"time"
)

// AddMedia handles post request for signup
func (h *Handler) AddMedia() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		file, _, err := c.Request.FormFile("file")
		defer file.Close()

		if err != nil {
			errresp.New(http.StatusBadRequest, false, err).
				SendResponse(c)
			return
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, file)

		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		randString, err := uuid.NewUUID()
		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		metadata := map[string]any{
			"id": randString,
		}

		h.ctrl.AddMedia(ctx, randString.String(), metadata, buf.Bytes())

		jsonresp.New(http.StatusAccepted, true, metadata, "file uploaded successfully").
			SendResponse(c)
	}
}
