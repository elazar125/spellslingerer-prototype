package controllers

import (
	"api/models"
	"api/rest"

	"github.com/gin-gonic/gin"
)

// GetTypes enumerates all card types
func GetTypes(c *gin.Context) {
	if types, err := models.AllCardTypes(); err != nil {
		rest.ServerError(c, err, "error loading card types")
	} else {
		rest.Ok(c, types)
	}
}

// GetSubTypes enumerates all subtypes
func GetSubTypes(c *gin.Context) {
	if subtypes, err := models.AllSubTypes(); err != nil {
		rest.ServerError(c, err, "error loading subtypes")
	} else {
		rest.Ok(c, subtypes)
	}
}

// GetColours enumerates all colours
func GetColours(c *gin.Context) {
	if colours, err := models.AllColours(); err != nil {
		rest.ServerError(c, err, "error loading colours")
	} else {
		rest.Ok(c, colours)
	}
}

// GetRarities enumerates all rarities
func GetRarities(c *gin.Context) {
	if rarities, err := models.AllRarities(); err != nil {
		rest.ServerError(c, err, "error loading rarities")
	} else {
		rest.Ok(c, rarities)
	}
}

// GetSets enumerates all sets
func GetSets(c *gin.Context) {
	if sets, err := models.AllSets(); err != nil {
		rest.ServerError(c, err, "error loading sets")
	} else {
		rest.Ok(c, sets)
	}
}

// GetArtists enumerates all artists
func GetArtists(c *gin.Context) {
	if artists, err := models.AllArtists(); err != nil {
		rest.ServerError(c, err, "error loading artists")
	} else {
		rest.Ok(c, artists)
	}
}

// GetMaterials enumerates all materials
func GetMaterials(c *gin.Context) {
	if materials, err := models.AllMaterials(); err != nil {
		rest.ServerError(c, err, "error loading materials")
	} else {
		rest.Ok(c, materials)
	}
}
