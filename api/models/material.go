package models

import "api/db"

// Material is used to craft cards
type Material struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllMaterials enumerates all materials
func AllMaterials() ([]Material, error) {
	var allMaterials []Material
	result := db.GlobalDB.Find(&allMaterials)
	return allMaterials, result.Error
}
