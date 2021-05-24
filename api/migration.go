package main

import (
	"api/db"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() error {
	if err := db.InitDatabase(); err != nil {
		return err
	} else if d, err := db.GlobalDB.DB(); err != nil {
		return err
	} else if driver, err := postgres.WithInstance(d, &postgres.Config{}); err != nil {
		return err
	} else if m, err := migrate.NewWithDatabaseInstance("file://./migrations", os.Getenv("DB_NAME"), driver); err != nil {
		return err
	} else {
		return m.Up()
	}
}
