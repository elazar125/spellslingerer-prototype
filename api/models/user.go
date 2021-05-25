package models

import (
	"api/db"
	"database/sql"
	"time"

	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

// User defines the user in db
type User struct {
	ID        uint
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	Name      string       `json:"name" binding:"required"`
	Email     string       `json:"email" binding:"required,email"`
	Password  string       `json:"password" binding:"required"`
}

type userMarshalJSON struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type userUnmarshalJSON struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// MarshalJSON overrides the default JSON marshaling to clear out the password prior to sending data
func (user User) MarshalJSON() ([]byte, error) {
	sanitized := userMarshalJSON{
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.Name,
		user.Email,
	}
	return json.Marshal(&sanitized)
}

// UnmarshalJSON overrides the default JSON unmarshaling to hash the password immediately before doing anything else with the data
func (user *User) UnmarshalJSON(bytes []byte) error {
	var unmarshal userUnmarshalJSON

	if err := json.Unmarshal(bytes, &unmarshal); err != nil {
		return err
	}

	user.ID = unmarshal.ID
	user.CreatedAt = unmarshal.CreatedAt
	user.UpdatedAt = unmarshal.UpdatedAt
	user.DeletedAt = unmarshal.DeletedAt
	user.Name = unmarshal.Name
	user.Email = unmarshal.Email
	user.Password = unmarshal.Password

	return user.hashPassword(user.Password)
}

// hashPassword encrypts user password
func (user *User) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	sql := `
		INSERT INTO public.users (created_at, updated_at, deleted_at, Name, Email, Password)
		VALUES (NOW(), NOW(), NULL, $1, $2, $3)
	`

	if _, err := db.GlobalSqlxDB.Exec(sql, user.Name, user.Email, user.Password); err != nil {
		return err
	}

	return db.GlobalSqlxDB.Get(user, "SELECT * FROM public.users U WHERE U.email = $1 LIMIT 1", user.Email)
}

// DeleteUserRecord deletes a user record in the database
func (user *User) DeleteUserRecord() error {
	_, err := db.GlobalSqlxDB.Exec("DELETE FROM public.users U WHERE U.email = $1", user.Email)
	return err
}

// UpdateUserRecord deletes a user record in the database
func (user *User) UpdateUserRecord(newData User) error {
	sql := `
		UPDATE public.users
		SET
			updated_at = NOW(),
			name = $1,
			email = $2,
			password = $3
		WHERE id = $4
	`

	if _, err := db.GlobalSqlxDB.Exec(sql, newData.Name, newData.Email, newData.Password, user.ID); err != nil {
		return err
	}

	return db.GlobalSqlxDB.Get(user, "SELECT * FROM public.users U WHERE U.email = $1 LIMIT 1", newData.Email)
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

// LookupByEmail finds the first (and theoretically only!) user with a given email, and loads it to the receiver
func (user *User) LookupByEmail(email string) error {
	return db.GlobalSqlxDB.Get(user, "SELECT * FROM public.users U WHERE U.email = $1 LIMIT 1", email)
}

// HasDuplicate ensures unique User properties (name, email) are not already in the DB
func (user *User) HasDuplicate() (bool, error) {
	var other User
	err := db.GlobalSqlxDB.Get(&other, "SELECT * FROM public.users U WHERE U.email = $1 OR U.name = $2", user.Email, user.Name)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
