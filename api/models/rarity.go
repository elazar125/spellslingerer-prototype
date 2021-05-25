package models

import "api/db"

// Rarity determines how likely a card is to be opened in a pack
type Rarity struct {
	ID   int64
	Name string
}

// AllRarities enumerates all rarities
func AllRarities() ([]Rarity, error) {
	var allRarities []Rarity
	err := db.GlobalSqlxDB.Select(&allRarities, "SELECT * FROM public.rarities")
	return allRarities, err
}
