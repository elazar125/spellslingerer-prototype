package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getTestCard(t *testing.T, cardName string) Card {
	set := Set{Name: cardName}
	err := set.Create()
	assert.NoError(t, err)
	rarity := Rarity{Name: cardName}
	err = rarity.Create()
	assert.NoError(t, err)
	cardType := CardType{Name: cardName}
	err = cardType.Create()
	assert.NoError(t, err)
	subType := SubType{Name: cardName}
	err = subType.Create()
	assert.NoError(t, err)
	artist := Artist{Name: cardName}
	err = artist.Create()
	assert.NoError(t, err)
	colour := Colour{Name: cardName}
	err = colour.Create()
	assert.NoError(t, err)
	material := Material{Name: cardName}
	err = material.Create()
	assert.NoError(t, err)

	return Card{
		Name:        cardName,
		Set:         set,
		Cost:        0,
		Power:       0,
		Health:      0,
		Ability:     "No ability",
		Chance:      0,
		Rarity:      rarity,
		CardType:    cardType,
		SubType:     subType,
		Artist:      artist,
		Colour:      colour,
		Legendary:   false,
		Collectible: false,
		Material:    material,
		Mox:         material,
	}
}

func cleanCardDependencies(t *testing.T, cardName string) {
	var set Set
	foundSet, err := set.Lookup(cardName)
	assert.NoError(t, err)
	err = foundSet.Delete()
	assert.NoError(t, err)
	var rarity Rarity
	foundRarity, err := rarity.Lookup(cardName)
	assert.NoError(t, err)
	err = foundRarity.Delete()
	assert.NoError(t, err)
	var cardType CardType
	foundCardType, err := cardType.Lookup(cardName)
	assert.NoError(t, err)
	err = foundCardType.Delete()
	assert.NoError(t, err)
	var subType SubType
	foundSubType, err := subType.Lookup(cardName)
	assert.NoError(t, err)
	err = foundSubType.Delete()
	assert.NoError(t, err)
	var artist Artist
	foundArtist, err := artist.Lookup(cardName)
	assert.NoError(t, err)
	err = foundArtist.Delete()
	assert.NoError(t, err)
	var colour Colour
	foundColour, err := colour.Lookup(cardName)
	assert.NoError(t, err)
	err = foundColour.Delete()
	assert.NoError(t, err)
	var material Material
	foundMaterial, err := material.Lookup(cardName)
	assert.NoError(t, err)
	err = foundMaterial.Delete()
	assert.NoError(t, err)
}

func TestCreateCard(t *testing.T) {
	var cardResult Card

	err := SetupDatabase(t)
	assert.NoError(t, err)

	card := getTestCard(t, "Test Create Card")

	err = card.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", card.Name).First(&cardResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", card.Name).Delete(&card)
	assert.NoError(t, result.Error)
	cleanCardDependencies(t, card.Name)

	assert.Equal(t, card.Name, cardResult.Name)
}

func TestDeleteCard(t *testing.T) {
	var cardResult Card

	err := SetupDatabase(t)
	assert.NoError(t, err)

	card := getTestCard(t, "Test Delete Card")

	err = card.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", card.Name).First(&cardResult)
	assert.NoError(t, result.Error)

	err = card.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", card.Name).First(&cardResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
	cleanCardDependencies(t, card.Name)
}

func TestLookupCard(t *testing.T) {
	err := SetupDatabase(t)
	assert.NoError(t, err)

	card := getTestCard(t, "Test Lookup Card")

	err = card.Create()
	assert.NoError(t, err)

	cardResult, err := card.Lookup(card.Name)
	assert.NoError(t, err)

	err = card.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, card.Name)

	assert.Equal(t, card.Name, cardResult.Name)
}

func TestUpdateCard(t *testing.T) {
	err := SetupDatabase(t)
	assert.NoError(t, err)

	card := getTestCard(t, "Test Update Card")

	err = card.Create()
	assert.NoError(t, err)

	firstCardResult, err := card.Lookup(card.Name)
	assert.NoError(t, err)

	card.Name = "New Name"
	err = card.Update()
	assert.NoError(t, err)

	secondCardResult, err := card.Lookup(card.Name)
	assert.NoError(t, err)

	_, err = card.Lookup("Test Update Card")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = card.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, "Test Update Card")

	assert.NotEqual(t, firstCardResult.Name, secondCardResult.Name)
	assert.NotEqual(t, firstCardResult.Name, card.Name)
	assert.Equal(t, secondCardResult.Name, card.Name)
}

func TestAllCards(t *testing.T) {
	err := SetupDatabase(t)
	assert.NoError(t, err)

	card1 := getTestCard(t, "Test All Cards 1")
	card2 := getTestCard(t, "Test All Cards 2")
	card3 := getTestCard(t, "Test All Cards 3")

	err = card1.Create()
	assert.NoError(t, err)
	err = card2.Create()
	assert.NoError(t, err)
	err = card3.Create()
	assert.NoError(t, err)

	allCards, err := card1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allCards))
	assert.Equal(t, card1.Name, allCards[0].Name)
	assert.Equal(t, card1.ID, allCards[0].ID)
	assert.Equal(t, card2.Name, allCards[1].Name)
	assert.Equal(t, card2.ID, allCards[1].ID)
	assert.Equal(t, card3.Name, allCards[2].Name)
	assert.Equal(t, card3.ID, allCards[2].ID)

	err = card1.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, card1.Name)
	err = card2.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, card2.Name)
	err = card3.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, card3.Name)
}

func TestUpsertCard(t *testing.T) {
	err := SetupDatabase(t)
	assert.NoError(t, err)

	card := getTestCard(t, "Test Upsert Card")

	_, err = card.Lookup("Test Upsert Card")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = card.Upsert()
	assert.NoError(t, err)

	firstCardResult, err := card.Lookup(card.Name)
	assert.NoError(t, err)

	card.Name = "New Name"
	err = card.Upsert()
	assert.NoError(t, err)

	secondCardResult, err := card.Lookup(card.Name)
	assert.NoError(t, err)

	_, err = card.Lookup("Test Upsert Card")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = card.Delete()
	assert.NoError(t, err)
	cleanCardDependencies(t, "Test Upsert Card")

	assert.NotEqual(t, firstCardResult.Name, secondCardResult.Name)
	assert.NotEqual(t, firstCardResult.Name, card.Name)
	assert.Equal(t, secondCardResult.Name, card.Name)
}
