package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateRarity(t *testing.T) {
	var rarityResult Rarity

	rarity := Rarity{
		Name: "Test Create Rarity",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = rarity.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", rarity.Name).First(&rarityResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", rarity.Name).Delete(&rarity)
	assert.NoError(t, result.Error)

	assert.Equal(t, rarity.Name, rarityResult.Name)
}

func TestDeleteRarity(t *testing.T) {
	var rarityResult Rarity

	rarity := Rarity{
		Name: "Test Delete Rarity",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = rarity.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", rarity.Name).First(&rarityResult)
	assert.NoError(t, result.Error)

	err = rarity.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", rarity.Name).First(&rarityResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupRarity(t *testing.T) {
	rarity := Rarity{
		Name: "Test Lookup Rarity",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = rarity.Create()
	assert.NoError(t, err)

	rarityResult, err := rarity.Lookup(rarity.Name)
	assert.NoError(t, err)

	err = rarity.Delete()
	assert.NoError(t, err)

	assert.Equal(t, rarity.Name, rarityResult.Name)
}

func TestUpdateRarity(t *testing.T) {
	rarity := Rarity{
		Name: "Test Update Rarity",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = rarity.Create()
	assert.NoError(t, err)

	firstRarityResult, err := rarity.Lookup(rarity.Name)
	assert.NoError(t, err)

	rarity.Name = "New Name"
	err = rarity.Update()
	assert.NoError(t, err)

	secondRarityResult, err := rarity.Lookup(rarity.Name)
	assert.NoError(t, err)

	_, err = rarity.Lookup("Test Update Rarity")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = rarity.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstRarityResult.Name, secondRarityResult.Name)
	assert.NotEqual(t, firstRarityResult.Name, rarity.Name)
	assert.Equal(t, secondRarityResult.Name, rarity.Name)
}

func TestAllRarities(t *testing.T) {
	rarity1 := Rarity{
		Name: "Test All Rarities 1",
	}
	rarity2 := Rarity{
		Name: "Test All Rarities 2",
	}
	rarity3 := Rarity{
		Name: "Test All Rarities 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = rarity1.Create()
	assert.NoError(t, err)
	err = rarity2.Create()
	assert.NoError(t, err)
	err = rarity3.Create()
	assert.NoError(t, err)

	allRarities, err := rarity1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allRarities))
	assert.Equal(t, rarity1.Name, allRarities[0].Name)
	assert.Equal(t, rarity1.ID, allRarities[0].ID)
	assert.Equal(t, rarity2.Name, allRarities[1].Name)
	assert.Equal(t, rarity2.ID, allRarities[1].ID)
	assert.Equal(t, rarity3.Name, allRarities[2].Name)
	assert.Equal(t, rarity3.ID, allRarities[2].ID)

	err = rarity1.Delete()
	assert.NoError(t, err)
	err = rarity2.Delete()
	assert.NoError(t, err)
	err = rarity3.Delete()
	assert.NoError(t, err)
}

func TestUpsertRarity(t *testing.T) {
	rarity := Rarity{
		Name: "Test Upsert Rarity",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = rarity.Lookup("Test Upsert Rarity")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = rarity.Upsert()
	assert.NoError(t, err)

	firstRarityResult, err := rarity.Lookup(rarity.Name)
	assert.NoError(t, err)

	rarity.Name = "New Name"
	err = rarity.Upsert()
	assert.NoError(t, err)

	secondRarityResult, err := rarity.Lookup(rarity.Name)
	assert.NoError(t, err)

	_, err = rarity.Lookup("Test Upsert Rarity")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = rarity.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstRarityResult.Name, secondRarityResult.Name)
	assert.NotEqual(t, firstRarityResult.Name, rarity.Name)
	assert.Equal(t, secondRarityResult.Name, rarity.Name)
}
