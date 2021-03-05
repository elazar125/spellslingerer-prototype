package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Material is what is used to craft new cards
type Material struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// All returns all materials in the database
func (material Material) All() ([]Material, error) {
	var allMaterials []Material
	result := db.GlobalDB.Find(&allMaterials)
	return allMaterials, result.Error
}

// Lookup finds a material in the database by a given name
func (material Material) Lookup(name string) (Material, error) {
	var foundMaterial Material
	result := db.GlobalDB.Where("name = ?", name).First(&foundMaterial)
	return foundMaterial, result.Error
}

// Create stores a record of the model in the database
func (material *Material) Create() error {
	result := db.GlobalDB.Create(&material)
	return result.Error
}

// Update updates a record of the model in the database
func (material *Material) Update() error {
	result := db.GlobalDB.Save(&material)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (material *Material) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&material)
	return result.Error
}

// Delete removes a record of the model from the database
func (material *Material) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", material.Name).Delete(&material)
	return result.Error
}
