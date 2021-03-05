package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateSet(t *testing.T) {
	var setResult Set

	set := Set{
		Name: "Test Create Set",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = set.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", set.Name).First(&setResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", set.Name).Delete(&set)
	assert.NoError(t, result.Error)

	assert.Equal(t, set.Name, setResult.Name)
}

func TestDeleteSet(t *testing.T) {
	var setResult Set

	set := Set{
		Name: "Test Delete Set",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = set.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", set.Name).First(&setResult)
	assert.NoError(t, result.Error)

	err = set.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", set.Name).First(&setResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupSet(t *testing.T) {
	set := Set{
		Name: "Test Lookup Set",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = set.Create()
	assert.NoError(t, err)

	setResult, err := set.Lookup(set.Name)
	assert.NoError(t, err)

	err = set.Delete()
	assert.NoError(t, err)

	assert.Equal(t, set.Name, setResult.Name)
}

func TestUpdateSet(t *testing.T) {
	set := Set{
		Name: "Test Update Set",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = set.Create()
	assert.NoError(t, err)

	firstSetResult, err := set.Lookup(set.Name)
	assert.NoError(t, err)

	set.Name = "New Name"
	err = set.Update()
	assert.NoError(t, err)

	secondSetResult, err := set.Lookup(set.Name)
	assert.NoError(t, err)

	_, err = set.Lookup("Test Update Set")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = set.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstSetResult.Name, secondSetResult.Name)
	assert.NotEqual(t, firstSetResult.Name, set.Name)
	assert.Equal(t, secondSetResult.Name, set.Name)
}

func TestAllSets(t *testing.T) {
	set1 := Set{
		Name: "Test All Sets 1",
	}
	set2 := Set{
		Name: "Test All Sets 2",
	}
	set3 := Set{
		Name: "Test All Sets 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = set1.Create()
	assert.NoError(t, err)
	err = set2.Create()
	assert.NoError(t, err)
	err = set3.Create()
	assert.NoError(t, err)

	allSets, err := set1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allSets))
	assert.Equal(t, set1.Name, allSets[0].Name)
	assert.Equal(t, set1.ID, allSets[0].ID)
	assert.Equal(t, set2.Name, allSets[1].Name)
	assert.Equal(t, set2.ID, allSets[1].ID)
	assert.Equal(t, set3.Name, allSets[2].Name)
	assert.Equal(t, set3.ID, allSets[2].ID)

	err = set1.Delete()
	assert.NoError(t, err)
	err = set2.Delete()
	assert.NoError(t, err)
	err = set3.Delete()
	assert.NoError(t, err)
}

func TestUpsertSet(t *testing.T) {
	set := Set{
		Name: "Test Upsert Set",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = set.Lookup("Test Upsert Set")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = set.Upsert()
	assert.NoError(t, err)

	firstSetResult, err := set.Lookup(set.Name)
	assert.NoError(t, err)

	set.Name = "New Name"
	err = set.Upsert()
	assert.NoError(t, err)

	secondSetResult, err := set.Lookup(set.Name)
	assert.NoError(t, err)

	_, err = set.Lookup("Test Upsert Set")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = set.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstSetResult.Name, secondSetResult.Name)
	assert.NotEqual(t, firstSetResult.Name, set.Name)
	assert.Equal(t, secondSetResult.Name, set.Name)
}
