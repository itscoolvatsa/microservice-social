package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"microservice/pkg/token"
	"net/http"
	"strings"
)

func CurrentUser(symmetricKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not Provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		accessToken := fields[1]
		maker, err := token.New(symmetricKey)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong while verifying"})
		}

		payload, err := maker.VerifyToken(accessToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		c.Set("access", payload)
		c.Next()
	}
}
