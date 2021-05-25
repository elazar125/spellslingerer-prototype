package models

import "api/db"

// Colour determines what decks a card can be used in
type Colour struct {
	ID   int64
	Name string
}

// AllColours enumerates all colours
func AllColours() ([]Colour, error) {
	var allColours []Colour
	err := db.GlobalSqlxDB.Select(&allColours, "SELECT * FROM public.colours")
	return allColours, err
}
