package auth

import (
	"api/rest"

	"github.com/gin-gonic/gin"
)

// EnsureAuthorized validates token and authorizes users
func EnsureAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if token, err := GetJwtTokenFromAuthHeader(authHeader); err != nil {
			rest.Forbidden(c, err, err.Error())
		} else if claims, err := token.validate(); err != nil {
			rest.Forbidden(c, err, err.Error())
		} else {
			c.Set("email", claims.Email)
			c.Next()
		}
	}
}
