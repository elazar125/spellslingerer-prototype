package models

import "api/db"

// Rarity determines how likely a card is to be opened in a pack
type Rarity struct {
	Id   int64  `gorm:"primary key"`
	Name string `gorm:"unique"`
}

// AllRarities enumerates all rarities
func AllRarities() ([]Rarity, error) {
	var allRarities []Rarity
	result := db.GlobalDB.Find(&allRarities)
	return allRarities, result.Error
}
