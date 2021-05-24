package models

import "api/db"

// Artist drew the card art
type Artist struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllArtists enumerates all artists
func AllArtists() ([]Artist, error) {
	var allArtists []Artist
	result := db.GlobalDB.Find(&allArtists)
	return allArtists, result.Error
}
