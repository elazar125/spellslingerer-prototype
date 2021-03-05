package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateColour(t *testing.T) {
	var colourResult Colour

	colour := Colour{
		Name: "Test Create Colour",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = colour.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", colour.Name).First(&colourResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", colour.Name).Delete(&colour)
	assert.NoError(t, result.Error)

	assert.Equal(t, colour.Name, colourResult.Name)
}

func TestDeleteColour(t *testing.T) {
	var colourResult Colour

	colour := Colour{
		Name: "Test Delete Colour",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = colour.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", colour.Name).First(&colourResult)
	assert.NoError(t, result.Error)

	err = colour.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", colour.Name).First(&colourResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupColour(t *testing.T) {
	colour := Colour{
		Name: "Test Lookup Colour",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = colour.Create()
	assert.NoError(t, err)

	colourResult, err := colour.Lookup(colour.Name)
	assert.NoError(t, err)

	err = colour.Delete()
	assert.NoError(t, err)

	assert.Equal(t, colour.Name, colourResult.Name)
}

func TestUpdateColour(t *testing.T) {
	colour := Colour{
		Name: "Test Update Colour",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = colour.Create()
	assert.NoError(t, err)

	firstColourResult, err := colour.Lookup(colour.Name)
	assert.NoError(t, err)

	colour.Name = "New Name"
	err = colour.Update()
	assert.NoError(t, err)

	secondColourResult, err := colour.Lookup(colour.Name)
	assert.NoError(t, err)

	_, err = colour.Lookup("Test Update Colour")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = colour.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstColourResult.Name, secondColourResult.Name)
	assert.NotEqual(t, firstColourResult.Name, colour.Name)
	assert.Equal(t, secondColourResult.Name, colour.Name)
}

func TestAllColours(t *testing.T) {
	colour1 := Colour{
		Name: "Test All Colours 1",
	}
	colour2 := Colour{
		Name: "Test All Colours 2",
	}
	colour3 := Colour{
		Name: "Test All Colours 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = colour1.Create()
	assert.NoError(t, err)
	err = colour2.Create()
	assert.NoError(t, err)
	err = colour3.Create()
	assert.NoError(t, err)

	allColours, err := colour1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allColours))
	assert.Equal(t, colour1.Name, allColours[0].Name)
	assert.Equal(t, colour1.ID, allColours[0].ID)
	assert.Equal(t, colour2.Name, allColours[1].Name)
	assert.Equal(t, colour2.ID, allColours[1].ID)
	assert.Equal(t, colour3.Name, allColours[2].Name)
	assert.Equal(t, colour3.ID, allColours[2].ID)

	err = colour1.Delete()
	assert.NoError(t, err)
	err = colour2.Delete()
	assert.NoError(t, err)
	err = colour3.Delete()
	assert.NoError(t, err)
}

func TestUpsertColour(t *testing.T) {
	colour := Colour{
		Name: "Test Upsert Colour",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = colour.Lookup("Test Upsert Colour")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = colour.Upsert()
	assert.NoError(t, err)

	firstColourResult, err := colour.Lookup(colour.Name)
	assert.NoError(t, err)

	colour.Name = "New Name"
	err = colour.Upsert()
	assert.NoError(t, err)

	secondColourResult, err := colour.Lookup(colour.Name)
	assert.NoError(t, err)

	_, err = colour.Lookup("Test Upsert Colour")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = colour.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstColourResult.Name, secondColourResult.Name)
	assert.NotEqual(t, firstColourResult.Name, colour.Name)
	assert.Equal(t, secondColourResult.Name, colour.Name)
}
