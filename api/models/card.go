package models

import (
	"gorm.io/gorm"
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

func (card Card) GetName() string {
	return card.Name
}

func (card Card) New(name string) (Card, error) {
	return Card{Name: name}, nil
}

func (card Card) GetID() int {
	return int(card.ID)
}
