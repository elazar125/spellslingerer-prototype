package models

import "api/db"

// CardType determines properties of the card and how it's played
type CardType struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllCardTypes enumerates all card types
func AllCardTypes() ([]CardType, error) {
	var allTypes []CardType
	result := db.GlobalDB.Find(&allTypes)
	return allTypes, result.Error
}
