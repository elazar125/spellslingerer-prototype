package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SubType drew the card art
type SubType struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all sub types in the database
func (subType SubType) All() ([]SubType, error) {
	var allSubTypes []SubType
	result := db.GlobalDB.Find(&allSubTypes)
	return allSubTypes, result.Error
}

// Lookup finds a sub type in the database by a given name
func (subType SubType) Lookup(name string) (SubType, error) {
	var foundSubType SubType
	result := db.GlobalDB.Where("name = ?", name).First(&foundSubType)
	return foundSubType, result.Error
}

// Create stores a record of the model in the database
func (subType *SubType) Create() error {
	result := db.GlobalDB.Create(&subType)
	return result.Error
}

// Update updates a record of the model in the database
func (subType *SubType) Update() error {
	result := db.GlobalDB.Save(&subType)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (subType *SubType) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&subType)
	return result.Error
}

// Delete removes a record of the model from the database
func (subType *SubType) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", subType.Name).Delete(&subType)
	return result.Error
}
