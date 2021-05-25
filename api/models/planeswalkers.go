package models

import "api/db"

// Planeswalker is used to craft cards
type Planeswalker struct {
	ID   int64
	Name string
}

// AllPlaneswalkers enumerates all planeswalkers
func AllPlaneswalkers() ([]Planeswalker, error) {
	var allPlaneswalkers []Planeswalker
	err := db.GlobalSqlxDB.Select(&allPlaneswalkers, "SELECT * FROM public.planeswalkers")
	return allPlaneswalkers, err
}
