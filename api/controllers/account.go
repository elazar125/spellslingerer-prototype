package controllers

import (
	"api/auth"
	"api/models"
	"api/rest"

	"database/sql"

	"github.com/gin-gonic/gin"
)

// Signup creates a user in db
func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		rest.BadRequest(c, err, "invalid json")
		return
	}

	if hasDuplicate, err := user.HasDuplicate(); hasDuplicate {
		rest.BadRequest(c, err, "user already exists")
		return
	} else if err != nil {
		rest.ServerError(c, err, "error creating user")
		return
	}

	if err := user.CreateUserRecord(); err != nil {
		rest.ServerError(c, err, "error creating user")
	} else {
		rest.Ok(c, user)
	}
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

	if err := user.LookupByEmail(payload.Email); err == sql.ErrNoRows {
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

// Logout logs users out
func Logout(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if token, err := auth.GetJwtTokenFromAuthHeader(authHeader); err != nil {
		rest.ServerError(c, err, "could not log out")
	} else if err := token.Blacklist(); err != nil {
		rest.ServerError(c, err, "could not log out")
	} else {
		rest.Ok(c, nil)
	}
}

// GetProfile logs users in
func GetProfile(c *gin.Context) {
	var user models.User

	if err := user.LookupByEmail(c.GetString("email")); err != nil {
		rest.ServerError(c, err, "could not get user profile")
	} else {
		rest.Ok(c, user)
	}
}

// EditProfile allows already authenticated users to update their name/email/password.
// Returns a new JWT token with te new email claim.
func EditProfile(c *gin.Context) {
	var user models.User
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		rest.BadRequest(c, err, "invalid json")
		return
	}

	if err := user.LookupByEmail(c.GetString("email")); err != nil {
		rest.ServerError(c, err, "could not get user profile")
		return
	}

	if err := user.UpdateUserRecord(newUser); err != nil {
		rest.ServerError(c, err, "error updating user")
		return
	}

	authHeader := c.Request.Header.Get("Authorization")
	if token, err := auth.GetJwtTokenFromAuthHeader(authHeader); err != nil {
		rest.ServerError(c, err, "could not log out")
		return
	} else if err := token.Blacklist(); err != nil {
		rest.ServerError(c, err, "could not log out")
		return
	}

	if signedToken, err := auth.GenerateJwtToken(user.Email); err != nil {
		rest.ServerError(c, err, "error signing token")
	} else {
		rest.Ok(c, signedToken)
	}
}

// DeleteProfile logs users in
func DeleteProfile(c *gin.Context) {
	var user models.User

	if err := user.LookupByEmail(c.GetString("email")); err != nil {
		rest.Unauthorized(c, err, "could not get user profile")
		return
	}

	if err := user.DeleteUserRecord(); err != nil {
		rest.ServerError(c, err, "could not delete user profile")
	} else {
		rest.Ok(c, user)
	}
}
