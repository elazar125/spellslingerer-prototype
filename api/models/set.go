package models

import "api/db"

// Set represents a group of cards published at the same time
type Set struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllSets enumerates all sets
func AllSets() ([]Set, error) {
	var allSets []Set
	result := db.GlobalDB.Find(&allSets)
	return allSets, result.Error
}
