package models

import (
	"api/db"

	"gorm.io/gorm"

	"testing"

	"github.com/stretchr/testify/assert"
)

func setupDatabase(t *testing.T) error {
	t.Log("setup database")

	err := db.InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = db.GlobalDB.AutoMigrate(&User{})
	assert.NoError(t, err)

	return err
}

func InsertUserInDatabase(t *testing.T, user User) func(t *testing.T) {
	t.Log("inserting user to database")

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	return func(t *testing.T) {
		user.DeleteUserRecord()
		t.Log("user \"Test User\" deleted from database")
	}
}

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)
	assert.NotEqual(t, "secret", user.Password)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	err := setupDatabase(t)
	assert.NoError(t, err)

	err = user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("email = ?", user.Email).First(&userResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("email = ?", user.Email).Delete(&user)
	assert.NoError(t, result.Error)

	assert.Equal(t, user.Name, userResult.Name)
	assert.Equal(t, user.Email, userResult.Email)
}

func TestDeleteUserRecord(t *testing.T) {
	var firstUserResult User
	var secondUserResult User

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	err := setupDatabase(t)
	assert.NoError(t, err)

	err = user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("email = ?", user.Email).First(&firstUserResult)
	assert.NoError(t, result.Error)

	err = user.DeleteUserRecord()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("email = ?", user.Email).First(&secondUserResult)
	assert.Error(t, result.Error)
	assert.Equal(t, result.Error, gorm.ErrRecordNotFound)
}

func TestCheckPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CheckPassword("secret")
	assert.NoError(t, err)

	err = user.CheckPassword("not the secret")
	assert.Error(t, err)
}

func TestLookupByEmail(t *testing.T) {
	var userResult User

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	err := setupDatabase(t)
	assert.NoError(t, err)

	teardown := InsertUserInDatabase(t, user)
	defer teardown(t)

	userResult.LookupByEmail(user.Email)

	assert.Equal(t, user.Name, userResult.Name)
	assert.Equal(t, user.Email, userResult.Email)
}

func TestHasDuplicate(t *testing.T) {
	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	err := setupDatabase(t)
	assert.NoError(t, err)

	hasDuplicate, err := user.HasDuplicate()
	assert.NoError(t, err)
	assert.False(t, hasDuplicate)

	teardown := InsertUserInDatabase(t, user)
	defer teardown(t)

	hasDuplicate, err = user.HasDuplicate()
	assert.NoError(t, err)
	assert.True(t, hasDuplicate)
}
