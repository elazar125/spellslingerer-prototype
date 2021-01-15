package models

import (
	"api/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User defines the user in db
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	result := db.GlobalDB.Create(&user)
	return result.Error
}

// DeleteUserRecord deletes a user record in the database
func (user *User) DeleteUserRecord() error {
	result := db.GlobalDB.Unscoped().Where("email = ?", user.Email).Delete(&user)
	return result.Error
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

// LookupByEmail finds the first (and theoretically only!) user with a given email, and loads it to the receiver
func (user *User) LookupByEmail(email string) error {
	result := db.GlobalDB.Where("email = ?", email).First(&user)
	return result.Error
}

// HasDuplicate ensures unique User properties (name, email) are not already in the DB
func (user *User) HasDuplicate() (bool, error) {
	var other User
	result := db.GlobalDB.Where("email = ? or name = ?", user.Email, user.Name).First(&other)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	} else if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}
