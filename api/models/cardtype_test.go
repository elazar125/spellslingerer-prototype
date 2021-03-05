package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateCardType(t *testing.T) {
	var cardTypeResult CardType

	cardType := CardType{
		Name: "Test Create CardType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = cardType.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", cardType.Name).First(&cardTypeResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", cardType.Name).Delete(&cardType)
	assert.NoError(t, result.Error)

	assert.Equal(t, cardType.Name, cardTypeResult.Name)
}

func TestDeleteCardType(t *testing.T) {
	var cardTypeResult CardType

	cardType := CardType{
		Name: "Test Delete CardType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = cardType.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", cardType.Name).First(&cardTypeResult)
	assert.NoError(t, result.Error)

	err = cardType.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", cardType.Name).First(&cardTypeResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupCardType(t *testing.T) {
	cardType := CardType{
		Name: "Test Lookup CardType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = cardType.Create()
	assert.NoError(t, err)

	cardTypeResult, err := cardType.Lookup(cardType.Name)
	assert.NoError(t, err)

	err = cardType.Delete()
	assert.NoError(t, err)

	assert.Equal(t, cardType.Name, cardTypeResult.Name)
}

func TestUpdateCardType(t *testing.T) {
	cardType := CardType{
		Name: "Test Update CardType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = cardType.Create()
	assert.NoError(t, err)

	firstCardTypeResult, err := cardType.Lookup(cardType.Name)
	assert.NoError(t, err)

	cardType.Name = "New Name"
	err = cardType.Update()
	assert.NoError(t, err)

	secondCardTypeResult, err := cardType.Lookup(cardType.Name)
	assert.NoError(t, err)

	_, err = cardType.Lookup("Test Update CardType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = cardType.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstCardTypeResult.Name, secondCardTypeResult.Name)
	assert.NotEqual(t, firstCardTypeResult.Name, cardType.Name)
	assert.Equal(t, secondCardTypeResult.Name, cardType.Name)
}

func TestAllCardTypes(t *testing.T) {
	cardType1 := CardType{
		Name: "Test All CardTypes 1",
	}
	cardType2 := CardType{
		Name: "Test All CardTypes 2",
	}
	cardType3 := CardType{
		Name: "Test All CardTypes 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = cardType1.Create()
	assert.NoError(t, err)
	err = cardType2.Create()
	assert.NoError(t, err)
	err = cardType3.Create()
	assert.NoError(t, err)

	allCardTypes, err := cardType1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allCardTypes))
	assert.Equal(t, cardType1.Name, allCardTypes[0].Name)
	assert.Equal(t, cardType1.ID, allCardTypes[0].ID)
	assert.Equal(t, cardType2.Name, allCardTypes[1].Name)
	assert.Equal(t, cardType2.ID, allCardTypes[1].ID)
	assert.Equal(t, cardType3.Name, allCardTypes[2].Name)
	assert.Equal(t, cardType3.ID, allCardTypes[2].ID)

	err = cardType1.Delete()
	assert.NoError(t, err)
	err = cardType2.Delete()
	assert.NoError(t, err)
	err = cardType3.Delete()
	assert.NoError(t, err)
}

func TestUpsertCardType(t *testing.T) {
	cardType := CardType{
		Name: "Test Upsert CardType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = cardType.Lookup("Test Upsert CardType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = cardType.Upsert()
	assert.NoError(t, err)

	firstCardTypeResult, err := cardType.Lookup(cardType.Name)
	assert.NoError(t, err)

	cardType.Name = "New Name"
	err = cardType.Upsert()
	assert.NoError(t, err)

	secondCardTypeResult, err := cardType.Lookup(cardType.Name)
	assert.NoError(t, err)

	_, err = cardType.Lookup("Test Upsert CardType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = cardType.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstCardTypeResult.Name, secondCardTypeResult.Name)
	assert.NotEqual(t, firstCardTypeResult.Name, cardType.Name)
	assert.Equal(t, secondCardTypeResult.Name, cardType.Name)
}
