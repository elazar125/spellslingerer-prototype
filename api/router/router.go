package router

import (
	"api/auth"
	"api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes route handling using the Gin framework
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("v1/")
	{
		// account handling, see account.go
		v1.POST("login", controllers.Login)
		v1.POST("signup", controllers.Signup)

		auth := v1.Group("/").Use(auth.EnsureAuthorized())
		{
			// account handling, see account.go
			auth.GET("profile", controllers.GetProfile)
			auth.DELETE("profile", controllers.DeleteProfile)
		}
	}

	return r
}
