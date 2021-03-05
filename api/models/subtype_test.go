package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateSubType(t *testing.T) {
	var subTypeResult SubType

	subType := SubType{
		Name: "Test Create SubType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = subType.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", subType.Name).First(&subTypeResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", subType.Name).Delete(&subType)
	assert.NoError(t, result.Error)

	assert.Equal(t, subType.Name, subTypeResult.Name)
}

func TestDeleteSubType(t *testing.T) {
	var subTypeResult SubType

	subType := SubType{
		Name: "Test Delete SubType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = subType.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", subType.Name).First(&subTypeResult)
	assert.NoError(t, result.Error)

	err = subType.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", subType.Name).First(&subTypeResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupSubType(t *testing.T) {
	subType := SubType{
		Name: "Test Lookup SubType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = subType.Create()
	assert.NoError(t, err)

	subTypeResult, err := subType.Lookup(subType.Name)
	assert.NoError(t, err)

	err = subType.Delete()
	assert.NoError(t, err)

	assert.Equal(t, subType.Name, subTypeResult.Name)
}

func TestUpdateSubType(t *testing.T) {
	subType := SubType{
		Name: "Test Update SubType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = subType.Create()
	assert.NoError(t, err)

	firstSubTypeResult, err := subType.Lookup(subType.Name)
	assert.NoError(t, err)

	subType.Name = "New Name"
	err = subType.Update()
	assert.NoError(t, err)

	secondSubTypeResult, err := subType.Lookup(subType.Name)
	assert.NoError(t, err)

	_, err = subType.Lookup("Test Update SubType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = subType.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstSubTypeResult.Name, secondSubTypeResult.Name)
	assert.NotEqual(t, firstSubTypeResult.Name, subType.Name)
	assert.Equal(t, secondSubTypeResult.Name, subType.Name)
}

func TestAllSubTypes(t *testing.T) {
	subType1 := SubType{
		Name: "Test All SubTypes 1",
	}
	subType2 := SubType{
		Name: "Test All SubTypes 2",
	}
	subType3 := SubType{
		Name: "Test All SubTypes 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = subType1.Create()
	assert.NoError(t, err)
	err = subType2.Create()
	assert.NoError(t, err)
	err = subType3.Create()
	assert.NoError(t, err)

	allSubTypes, err := subType1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allSubTypes))
	assert.Equal(t, subType1.Name, allSubTypes[0].Name)
	assert.Equal(t, subType1.ID, allSubTypes[0].ID)
	assert.Equal(t, subType2.Name, allSubTypes[1].Name)
	assert.Equal(t, subType2.ID, allSubTypes[1].ID)
	assert.Equal(t, subType3.Name, allSubTypes[2].Name)
	assert.Equal(t, subType3.ID, allSubTypes[2].ID)

	err = subType1.Delete()
	assert.NoError(t, err)
	err = subType2.Delete()
	assert.NoError(t, err)
	err = subType3.Delete()
	assert.NoError(t, err)
}

func TestUpsertSubType(t *testing.T) {
	subType := SubType{
		Name: "Test Upsert SubType",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = subType.Lookup("Test Upsert SubType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = subType.Upsert()
	assert.NoError(t, err)

	firstSubTypeResult, err := subType.Lookup(subType.Name)
	assert.NoError(t, err)

	subType.Name = "New Name"
	err = subType.Upsert()
	assert.NoError(t, err)

	secondSubTypeResult, err := subType.Lookup(subType.Name)
	assert.NoError(t, err)

	_, err = subType.Lookup("Test Upsert SubType")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = subType.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstSubTypeResult.Name, secondSubTypeResult.Name)
	assert.NotEqual(t, firstSubTypeResult.Name, subType.Name)
	assert.Equal(t, secondSubTypeResult.Name, subType.Name)
}
