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

		v1.GET("types", controllers.GetTypes)
		v1.GET("subtypes", controllers.GetSubTypes)
		v1.GET("colours", controllers.GetColours)
		v1.GET("sets", controllers.GetSets)
		v1.GET("rarities", controllers.GetRarities)
		v1.GET("artists", controllers.GetArtists)
		v1.GET("materials", controllers.GetMaterials)
		// v1.GET("planeswalkers", controllers.GetPlaneswalkers) TODO: figure out schema for this
		v1.GET("cards", controllers.GetCards)

		auth := v1.Group("/").Use(auth.EnsureAuthorized())
		{
			// account handling, see account.go
			auth.POST("logout", controllers.Logout)
			auth.GET("profile", controllers.GetProfile)
			auth.PUT("profile", controllers.EditProfile)
			auth.DELETE("profile", controllers.DeleteProfile)
		}
	}

	return r
}
