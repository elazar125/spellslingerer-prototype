package models

import "api/db"

// SubType is a short description of what kind of creature is on the card
type SubType struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllSubTypes enumerates all subtypes
func AllSubTypes() ([]SubType, error) {
	var allSubTypes []SubType
	result := db.GlobalDB.Find(&allSubTypes)
	return allSubTypes, result.Error
}
