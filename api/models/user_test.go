package models

import (
	"api/auth"
	"api/db"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// setupDBForUsers sets up the database for use in unit tests
// TODO: mock the database
func setupDBForUsers(t *testing.T) error {
	t.Log("setup database")

	err := db.InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = db.GlobalDB.AutoMigrate(User{})
	assert.NoError(t, err)
	err = db.GlobalDB.AutoMigrate(auth.SignedToken{})
	assert.NoError(t, err)

	return err
}

func insertUserInDatabase(t *testing.T, user User) func(t *testing.T) {
	t.Log("inserting user to database")

	err := user.CreateUserRecord()
	assert.NoError(t, err)

	return func(t *testing.T) {
		user.DeleteUserRecord()
		t.Log("user \"Test User\" deleted from database")
	}
}

func TestMarshalJSON(t *testing.T) {
	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	expected := userMarshalJSON{
		user.Model,
		user.Name,
		user.Email,
		// Note Password is not included
	}

	testBytes, err := json.Marshal(user)
	assert.NoError(t, err)

	compareBytes, err := json.Marshal(expected)
	assert.NoError(t, err)
	assert.Equal(t, compareBytes, testBytes)
}

func TestUnmarshalJSON(t *testing.T) {
	var testUser User

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	toMarshal := userUnmarshalJSON{
		user.Model,
		user.Name,
		user.Email,
		user.Password,
	}

	// Marshal a copy to avoid the password being stripped out
	testBytes, err := json.Marshal(toMarshal)
	assert.NoError(t, err)

	err = json.Unmarshal(testBytes, &testUser)
	assert.NoError(t, err)

	assert.Equal(t, user.Model, testUser.Model)
	assert.Equal(t, user.Name, testUser.Name)
	assert.Equal(t, user.Email, testUser.Email)
	assert.NotEqual(t, user.Password, testUser.Password)
	err = testUser.CheckPassword(user.Password)
	assert.NoError(t, err)
}

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.hashPassword(user.Password)
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

	err := setupDBForUsers(t)
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

	err := setupDBForUsers(t)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("email = ?", user.Email).First(&firstUserResult)
	assert.NoError(t, result.Error)

	err = user.DeleteUserRecord()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("email = ?", user.Email).First(&secondUserResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestUpdateUserRecord(t *testing.T) {
	var userResult User

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "secret",
	}

	err := setupDBForUsers(t)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	newUser := User{
		Name:     "New Name",
		Email:    "new@email.com",
		Password: "new secret",
	}

	err = user.UpdateUserRecord(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.Password, user.Password)

	err = userResult.LookupByEmail(newUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, userResult.Name)
	assert.Equal(t, newUser.Email, userResult.Email)
	assert.Equal(t, newUser.Password, userResult.Password)

	err = user.DeleteUserRecord()
	assert.NoError(t, err)
}

func TestCheckPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.hashPassword(user.Password)
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

	err := setupDBForUsers(t)
	assert.NoError(t, err)

	teardown := insertUserInDatabase(t, user)
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

	err := setupDBForUsers(t)
	assert.NoError(t, err)

	hasDuplicate, err := user.HasDuplicate()
	assert.NoError(t, err)
	assert.False(t, hasDuplicate)

	teardown := insertUserInDatabase(t, user)
	defer teardown(t)

	hasDuplicate, err = user.HasDuplicate()
	assert.NoError(t, err)
	assert.True(t, hasDuplicate)
}
