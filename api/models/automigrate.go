package models

import "api/db"

// AutoMigrateModels ensures the models exist in the database schema
func AutoMigrateModels() error {
	if err := db.GlobalDB.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Set{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Artist{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Material{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&CardType{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&SubType{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Rarity{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Colour{}); err != nil {
		return err
	}
	if err := db.GlobalDB.AutoMigrate(&Card{}); err != nil {
		return err
	}

	return nil
}
