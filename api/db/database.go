package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GlobalDB is a global database connection, call InitDatabase before using
var GlobalDB *gorm.DB

// InitDatabase initializes the global database connection GlobalDB
func InitDatabase() (err error) {
	conn := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v database=%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	GlobalDB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	return
}
