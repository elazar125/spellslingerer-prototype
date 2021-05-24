package controllers

import (
	"api/models"
	"api/rest"

	"github.com/gin-gonic/gin"
)

// GetPlaneswalkers enumerates all planeswalkers
func GetPlaneswalkers(c *gin.Context) {
	if planeswalkers, err := models.AllPlaneswalkers(); err != nil {
		rest.ServerError(c, err, "error loading planeswalkers")
	} else {
		rest.Ok(c, planeswalkers)
	}
}
