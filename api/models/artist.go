package models

import "api/db"

// Artist drew the card art
type Artist struct {
	ID   int64
	Name string
}

// AllArtists enumerates all artists
func AllArtists() ([]Artist, error) {
	var allArtists []Artist
	err := db.GlobalSqlxDB.Select(&allArtists, "SELECT * FROM public.artists")
	return allArtists, err
}
