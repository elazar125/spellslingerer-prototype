package models

import "api/db"

// Planeswalker represents a deck archetype
type Planeswalker struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllPlaneswalkers enumerates all planeswalkers
func AllPlaneswalkers() ([]Planeswalker, error) {
	var allPlaneswalkers []Planeswalker
	result := db.GlobalDB.Find(&allPlaneswalkers)
	return allPlaneswalkers, result.Error
}
