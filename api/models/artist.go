package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Artist drew the card art
type Artist struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all artists in the database
func (artist Artist) All() ([]Artist, error) {
	var allArtists []Artist
	result := db.GlobalDB.Find(&allArtists)
	return allArtists, result.Error
}

// Lookup finds an artist in the database by a given name
func (artist Artist) Lookup(name string) (Artist, error) {
	var foundArtist Artist
	result := db.GlobalDB.Where("name = ?", name).First(&foundArtist)
	return foundArtist, result.Error
}

// Create stores a record of the model in the database
func (artist *Artist) Create() error {
	result := db.GlobalDB.Create(&artist)
	return result.Error
}

// Update updates a record of the model in the database
func (artist *Artist) Update() error {
	result := db.GlobalDB.Save(&artist)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (artist *Artist) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&artist)
	return result.Error
}

// Delete removes a record of the model from the database
func (artist *Artist) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", artist.Name).Delete(&artist)
	return result.Error
}
