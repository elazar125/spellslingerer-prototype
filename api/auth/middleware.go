package auth

import (
	"api/rest"

	"strings"

	"github.com/gin-gonic/gin"
)

// EnsureAuthorized validates token and authorizes users
func EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(bearerToken, "Bearer ")
		claims, err := ValidateJwtToken(token)

		if err != nil {
			rest.Forbidden(c, err, err.Error())
			return
		}

		c.Set("email", claims.Email)

		c.Next()
	}
}
