package controllers

import (
	"api/models"
	"api/rest"

	"github.com/gin-gonic/gin"
)

// GetCards enumerates all cards
func GetCards(c *gin.Context) {
	if cards, err := models.AllCards(); err != nil {
		rest.ServerError(c, err, "error loading cards")
	} else {
		rest.Ok(c, cards)
	}
}
