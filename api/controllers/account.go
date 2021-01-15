package controllers

import (
	"api/auth"
	"api/models"
	"api/rest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Signup creates a user in db
func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		rest.BadRequest(c, err, "invalid json")
		return
	}

	if hasDuplicate, err := user.HasDuplicate(); hasDuplicate || err != nil {
		rest.BadRequest(c, err, "user already exists")
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		rest.ServerError(c, err, "error hashing password")
		return
	}

	if err := user.CreateUserRecord(); err != nil {
		rest.ServerError(c, err, "error creating user")
		return
	}

	rest.Ok(c, user)
}

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login logs users in
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		rest.BadRequest(c, err, "invalid json")
		return
	}

	if err := user.LookupByEmail(payload.Email); err == gorm.ErrRecordNotFound {
		rest.Unauthorized(c, err, "invalid user credentials")
		return
	} else if err != nil {
		rest.ServerError(c, err, "could not get user profile")
		return
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		rest.Unauthorized(c, err, "invalid user credentials")
		return
	}

	if signedToken, err := auth.GenerateJwtToken(user.Email); err != nil {
		rest.ServerError(c, err, "error signing token")
	} else {
		rest.Ok(c, signedToken)
	}
}

// MyData logs users in
func MyData(c *gin.Context) {
	var user models.User

	if err := user.LookupByEmail(c.GetString("email")); err == gorm.ErrRecordNotFound {
		rest.Unauthorized(c, err, "invalid user credentials")
		return
	} else if err != nil {
		rest.ServerError(c, err, "could not get user profile")
		return
	}

	user.Password = ""

	rest.Ok(c, user)
}

// DeleteMyData logs users in
func DeleteMyData(c *gin.Context) {
	var user models.User

	if err := user.LookupByEmail(c.GetString("email")); err == gorm.ErrRecordNotFound {
		rest.Unauthorized(c, err, "invalid user credentials")
		return
	} else if err != nil {
		rest.ServerError(c, err, "could not get user profile")
		return
	}

	if err := user.DeleteUserRecord(); err != nil {
		rest.ServerError(c, err, "could not delete user profile")
	} else {
		rest.Ok(c, user)
	}
}
