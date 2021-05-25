package models

import "api/db"

// CardType determines properties of the card and how it's played
type CardType struct {
	ID   int64
	Name string
}

// AllCardTypes enumerates all card types
func AllCardTypes() ([]CardType, error) {
	var allTypes []CardType
	err := db.GlobalSqlxDB.Select(&allTypes, "SELECT * FROM public.card_types")
	return allTypes, err
}
