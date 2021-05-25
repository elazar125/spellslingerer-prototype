package models

import "api/db"

// Set represents a group of cards published at the same time
type Set struct {
	ID   int64
	Name string
}

// AllSets enumerates all sets
func AllSets() ([]Set, error) {
	var allSets []Set
	err := db.GlobalSqlxDB.Select(&allSets, "SELECT * FROM public.sets")
	return allSets, err
}
