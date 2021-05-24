package models

import "api/db"

// Colour determines what decks a card can be used in
type Colour struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllColours enumerates all colours
func AllColours() ([]Colour, error) {
	var allColours []Colour
	result := db.GlobalDB.Find(&allColours)
	return allColours, result.Error
}
