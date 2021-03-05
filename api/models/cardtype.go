package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CardType determines properties of the card and how it's played
type CardType struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all card types in the database
func (cardType CardType) All() ([]CardType, error) {
	var allCardTypes []CardType
	result := db.GlobalDB.Find(&allCardTypes)
	return allCardTypes, result.Error
}

// Lookup finds a card type in the database by a given name
func (cardType CardType) Lookup(name string) (CardType, error) {
	var foundCardType CardType
	result := db.GlobalDB.Where("name = ?", name).First(&foundCardType)
	return foundCardType, result.Error
}

// Create stores a record of the model in the database
func (cardType *CardType) Create() error {
	result := db.GlobalDB.Create(&cardType)
	return result.Error
}

// Update updates a record of the model in the database
func (cardType *CardType) Update() error {
	result := db.GlobalDB.Save(&cardType)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (cardType *CardType) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&cardType)
	return result.Error
}

// Delete removes a record of the model from the database
func (cardType *CardType) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", cardType.Name).Delete(&cardType)
	return result.Error
}
