package middleware

import (
	"github.com/gin-gonic/gin"
	"microservice/pkg/errors"
	errresp "microservice/pkg/response/error"
	"microservice/pkg/token"
	"net/http"
	"strings"
)

func CurrentUser(symmetricKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if len(authorizationHeader) == 0 {
			resp := errresp.New(http.StatusUnauthorized, false, errors.ErrNoAuthHeader)
			resp.SendResponse(c)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			resp := errresp.New(http.StatusUnauthorized, false, errors.ErrInvalidAuthHeader)
			resp.SendResponse(c)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			resp := errresp.New(http.StatusUnauthorized, false, errors.ErrInvalidAuthData)
			resp.SendResponse(c)
			return
		}

		accessToken := fields[1]
		maker, err := token.New(symmetricKey)

		if err != nil {
			resp := errresp.New(http.StatusUnauthorized, false, errors.ErrInternalServer)
			resp.SendResponse(c)
			return
		}

		payload, err := maker.VerifyToken(accessToken)

		if err != nil {
			resp := errresp.New(http.StatusUnauthorized, false, errors.ErrInvalidAuthData)
			resp.SendResponse(c)
			return
		}

		c.Set("access", payload)
		c.Next()
	}
}
