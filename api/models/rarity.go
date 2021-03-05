package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Rarity determines how rare the card is
type Rarity struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all rarities in the database
func (rarity Rarity) All() ([]Rarity, error) {
	var allRarities []Rarity
	result := db.GlobalDB.Find(&allRarities)
	return allRarities, result.Error
}

// Lookup finds a rarity in the database by a given name
func (rarity Rarity) Lookup(name string) (Rarity, error) {
	var foundRarity Rarity
	result := db.GlobalDB.Where("name = ?", name).First(&foundRarity)
	return foundRarity, result.Error
}

// Create stores a record of the model in the database
func (rarity *Rarity) Create() error {
	result := db.GlobalDB.Create(&rarity)
	return result.Error
}

// Update updates a record of the model in the database
func (rarity *Rarity) Update() error {
	result := db.GlobalDB.Save(&rarity)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (rarity *Rarity) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&rarity)
	return result.Error
}

// Delete removes a record of the model from the database
func (rarity *Rarity) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", rarity.Name).Delete(&rarity)
	return result.Error
}
