package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateMaterial(t *testing.T) {
	var materialResult Material

	material := Material{
		Name: "Test Create Material",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = material.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", material.Name).First(&materialResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", material.Name).Delete(&material)
	assert.NoError(t, result.Error)

	assert.Equal(t, material.Name, materialResult.Name)
}

func TestDeleteMaterial(t *testing.T) {
	var materialResult Material

	material := Material{
		Name: "Test Delete Material",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = material.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", material.Name).First(&materialResult)
	assert.NoError(t, result.Error)

	err = material.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", material.Name).First(&materialResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupMaterial(t *testing.T) {
	material := Material{
		Name: "Test Lookup Material",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = material.Create()
	assert.NoError(t, err)

	materialResult, err := material.Lookup(material.Name)
	assert.NoError(t, err)

	err = material.Delete()
	assert.NoError(t, err)

	assert.Equal(t, material.Name, materialResult.Name)
}

func TestUpdateMaterial(t *testing.T) {
	material := Material{
		Name: "Test Update Material",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = material.Create()
	assert.NoError(t, err)

	firstMaterialResult, err := material.Lookup(material.Name)
	assert.NoError(t, err)

	material.Name = "New Name"
	err = material.Update()
	assert.NoError(t, err)

	secondMaterialResult, err := material.Lookup(material.Name)
	assert.NoError(t, err)

	_, err = material.Lookup("Test Update Material")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = material.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstMaterialResult.Name, secondMaterialResult.Name)
	assert.NotEqual(t, firstMaterialResult.Name, material.Name)
	assert.Equal(t, secondMaterialResult.Name, material.Name)
}

func TestAllMaterials(t *testing.T) {
	material1 := Material{
		Name: "Test All Materials 1",
	}
	material2 := Material{
		Name: "Test All Materials 2",
	}
	material3 := Material{
		Name: "Test All Materials 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = material1.Create()
	assert.NoError(t, err)
	err = material2.Create()
	assert.NoError(t, err)
	err = material3.Create()
	assert.NoError(t, err)

	allMaterials, err := material1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allMaterials))
	assert.Equal(t, material1.Name, allMaterials[0].Name)
	assert.Equal(t, material1.ID, allMaterials[0].ID)
	assert.Equal(t, material2.Name, allMaterials[1].Name)
	assert.Equal(t, material2.ID, allMaterials[1].ID)
	assert.Equal(t, material3.Name, allMaterials[2].Name)
	assert.Equal(t, material3.ID, allMaterials[2].ID)

	err = material1.Delete()
	assert.NoError(t, err)
	err = material2.Delete()
	assert.NoError(t, err)
	err = material3.Delete()
	assert.NoError(t, err)
}

func TestUpsertMaterial(t *testing.T) {
	material := Material{
		Name: "Test Upsert Material",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = material.Lookup("Test Upsert Material")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = material.Upsert()
	assert.NoError(t, err)

	firstMaterialResult, err := material.Lookup(material.Name)
	assert.NoError(t, err)

	material.Name = "New Name"
	err = material.Upsert()
	assert.NoError(t, err)

	secondMaterialResult, err := material.Lookup(material.Name)
	assert.NoError(t, err)

	_, err = material.Lookup("Test Upsert Material")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = material.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstMaterialResult.Name, secondMaterialResult.Name)
	assert.NotEqual(t, firstMaterialResult.Name, material.Name)
	assert.Equal(t, secondMaterialResult.Name, material.Name)
}
