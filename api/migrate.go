package main

import (
	"api/auth"
	"api/db"
	"api/models"
)

func setupAndMigrateDatabase() error {
	if err := db.InitDatabase(); err != nil {
		return err
	}

	if err := db.GlobalDB.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&auth.SignedToken{}); err != nil {
		return err
	}

	return nil
}
