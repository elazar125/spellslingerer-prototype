package models

import (
	"api/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Card defines the properties of a card in Spellslingers
type Card struct {
	gorm.Model
	Name        string `gorm:"unique"`
	SetID       int
	Set         Set `gorm:"foreignkey:SetID"`
	Cost        int
	Power       int
	Health      int
	Ability     string
	Chance      int
	RarityID    int
	Rarity      Rarity `gorm:"foreignkey:RarityID"`
	CardTypeID  int
	CardType    CardType `gorm:"foreignkey:CardTypeID"`
	SubTypeID   int
	SubType     SubType `gorm:"foreignkey:SubTypeID"`
	ArtistID    int
	Artist      Artist `gorm:"foreignkey:ArtistID"`
	ColourID    int
	Colour      Colour `gorm:"foreignkey:ColourID"`
	Legendary   bool
	Collectible bool
	MaterialID  int
	Material    Material `gorm:"foreignkey:MaterialID"`
	MoxID       int
	Mox         Material `gorm:"foreignkey:MoxID"`
}

// All returns all cards in the database
func (card Card) All() ([]Card, error) {
	var allCards []Card
	result := db.GlobalDB.Find(&allCards)
	return allCards, result.Error
}

// Lookup finds a card in the database by a given name
func (card Card) Lookup(name string) (Card, error) {
	var foundCard Card
	result := db.GlobalDB.Where("name = ?", name).First(&foundCard)
	return foundCard, result.Error
}

// Create stores a record of the model in the database
func (card *Card) Create() error {
	result := db.GlobalDB.Create(&card)
	return result.Error
}

// Update updates a record of the model in the database
func (card *Card) Update() error {
	result := db.GlobalDB.Save(&card)
	return result.Error
}

// Upsert stores a record of the model in the database if it doesn't exist,
// otherwise updates it
func (card *Card) Upsert() error {
	result := db.GlobalDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&card)
	return result.Error
}

// Delete removes a record of the model from the database
func (card *Card) Delete() error {
	result := db.GlobalDB.Unscoped().Where("name = ?", card.Name).Delete(&card)
	return result.Error
}
