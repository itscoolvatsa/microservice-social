package http

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	jsonresp "microservice/pkg/response/json"
	"microservice/pkg/token"
	"microservice/pkg/util"
	"microservice/posts/pkg/model"
	"net/http"
	"time"
)

// AddPost handles post request for signup
func (h *Handler) AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		payload := c.MustGet("access").(*token.Payload)

		file, _, err := c.Request.FormFile("file")

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
		}
		name := util.RandomString(10)

		fid, err := h.ctrl.AddMedia(name, buf.Bytes())

		if err != nil {
			errresp.New(http.StatusInternalServerError, false, errors.ErrInternalServer).
				SendResponse(c)
		}

		var post model.Post

		caption := c.Request.FormValue("caption")
		defer cancel()

		post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.ImageId = fid
		post.Caption = caption
		post.UserId = payload.UserId

		//	saving the suer to the database
		err = h.ctrl.AddPost(ctx, &post)

		if err != nil {
			errresp.New(http.StatusBadRequest, false, errors.ErrInternalServer).
				SendResponse(c)
			return
		}

		jsonresp.New(http.StatusAccepted, true, post, "post added successfully").
			SendResponse(c)
	}
}
