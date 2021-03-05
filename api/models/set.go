package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Set determines properties of the card and how it's played
type Set struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all sets in the database
func (set Set) All() ([]Set, error) {
	var allSet []Set
	result := db.GlobalDB.Find(&allSet)
	return allSet, result.Error
}

// Lookup finds a card type in the database by a given name
func (set Set) Lookup(name string) (Set, error) {
	var foundSet Set
	result := db.GlobalDB.Where("name = ?", name).First(&foundSet)
	return foundSet, result.Error
}

// Create stores a record of the model in the database
func (set *Set) Create() error {
	result := db.GlobalDB.Create(&set)
	return result.Error
}

// Update updates a record of the model in the database
func (set *Set) Update() error {
	result := db.GlobalDB.Save(&set)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (set *Set) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&set)
	return result.Error
}

// Delete removes a record of the model from the database
func (set *Set) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", set.Name).Delete(&set)
	return result.Error
}
