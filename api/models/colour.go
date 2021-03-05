package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Colour determines which decks can include a card
type Colour struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all colours in the database
func (colour Colour) All() ([]Colour, error) {
	var allColours []Colour
	result := db.GlobalDB.Find(&allColours)
	return allColours, result.Error
}

// Lookup finds a colour in the database by a given name
func (colour Colour) Lookup(name string) (Colour, error) {
	var foundColour Colour
	result := db.GlobalDB.Where("name = ?", name).First(&foundColour)
	return foundColour, result.Error
}

// Create stores a record of the model in the database
func (colour *Colour) Create() error {
	result := db.GlobalDB.Create(&colour)
	return result.Error
}

// Update updates a record of the model in the database
func (colour *Colour) Update() error {
	result := db.GlobalDB.Save(&colour)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (colour *Colour) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&colour)
	return result.Error
}

// Delete removes a record of the model from the database
func (colour *Colour) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", colour.Name).Delete(&colour)
	return result.Error
}
