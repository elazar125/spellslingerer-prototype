package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Ok returns an Ok response after logging any relevant info
func Ok(c *gin.Context, body interface{}) {
	c.JSON(200, body)
}

// BadRequest returns a BadRequest response after logging any relevant info
func BadRequest(c *gin.Context, err error, message string) {
	log.Println(err)

	c.JSON(400, gin.H{
		"msg": message,
	})
	c.Abort()
}

// Unauthorized returns an Unauthorized response after logging any relevant info
func Unauthorized(c *gin.Context, err error, message string) {
	log.Println(err)

	c.JSON(401, gin.H{
		"msg": message,
	})
	c.Abort()
}

// Forbidden returns a Forbidden response after logging any relevant info
func Forbidden(c *gin.Context, err error, message string) {
	log.Println(err)

	c.JSON(403, gin.H{
		"msg": message,
	})
	c.Abort()
}

// NotFound returns a NotFound response after logging any relevant info
func NotFound(c *gin.Context, err error, message string) {
	log.Println(err)

	c.JSON(404, gin.H{
		"msg": message,
	})
	c.Abort()
}

// ServerError returns a ServerError response after logging any relevant info
func ServerError(c *gin.Context, err error, message string) {
	log.Println(err)

	c.JSON(500, gin.H{
		"msg": message,
	})
	c.Abort()
}
