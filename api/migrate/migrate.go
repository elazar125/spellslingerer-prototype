package migrate

import (
	"api/auth"
	"api/db"
	"api/models"
)

// SetupAndMigrateDatabase initializes the database and runs all migrations needed
func SetupAndMigrateDatabase() error {
	if err := db.InitDatabase(); err != nil {
		return err
	}

	if err := models.AutoMigrateModels(); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&auth.SignedToken{}); err != nil {
		return err
	}

	return nil
}
