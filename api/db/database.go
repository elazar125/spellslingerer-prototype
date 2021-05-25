package db

import (
	"fmt"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

// GlobalSqlxDB is a global database connection, call InitDatabase before using
var GlobalSqlxDB *sqlx.DB

// InitDatabase initializes the global database connection GlobalSqlxDB
func InitDatabase() (err error) {
	conn := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v database=%v sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	GlobalSqlxDB, err = sqlx.Connect("postgres", conn)

	return
}

func Migrate() error {
	if err := InitDatabase(); err != nil {
		return err
	} else if driver, err := postgres.WithInstance(GlobalSqlxDB.DB, &postgres.Config{}); err != nil {
		return err
	} else if m, err := migrate.NewWithDatabaseInstance("file://./migrations", os.Getenv("DB_NAME"), driver); err != nil {
		return err
	} else if err = m.Up(); err != migrate.ErrNoChange {
		return err
	} else {
		return nil
	}
}
