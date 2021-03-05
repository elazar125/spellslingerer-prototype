package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateArtist(t *testing.T) {
	var artistResult Artist

	artist := Artist{
		Name: "Test Create Artist",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = artist.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", artist.Name).First(&artistResult)
	assert.NoError(t, result.Error)

	result = db.GlobalDB.Unscoped().Where("name = ?", artist.Name).Delete(&artist)
	assert.NoError(t, result.Error)

	assert.Equal(t, artist.Name, artistResult.Name)
}

func TestDeleteArtist(t *testing.T) {
	var artistResult Artist

	artist := Artist{
		Name: "Test Delete Artist",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = artist.Create()
	assert.NoError(t, err)

	result := db.GlobalDB.Where("name = ?", artist.Name).First(&artistResult)
	assert.NoError(t, result.Error)

	err = artist.Delete()
	assert.NoError(t, err)

	result = db.GlobalDB.Where("name = ?", artist.Name).First(&artistResult)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}

func TestLookupArtist(t *testing.T) {
	artist := Artist{
		Name: "Test Lookup Artist",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = artist.Create()
	assert.NoError(t, err)

	artistResult, err := artist.Lookup(artist.Name)
	assert.NoError(t, err)

	err = artist.Delete()
	assert.NoError(t, err)

	assert.Equal(t, artist.Name, artistResult.Name)
}

func TestUpdateArtist(t *testing.T) {
	artist := Artist{
		Name: "Test Update Artist",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = artist.Create()
	assert.NoError(t, err)

	firstArtistResult, err := artist.Lookup(artist.Name)
	assert.NoError(t, err)

	artist.Name = "New Name"
	err = artist.Update()
	assert.NoError(t, err)

	secondArtistResult, err := artist.Lookup(artist.Name)
	assert.NoError(t, err)

	_, err = artist.Lookup("Test Update Artist")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = artist.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstArtistResult.Name, secondArtistResult.Name)
	assert.NotEqual(t, firstArtistResult.Name, artist.Name)
	assert.Equal(t, secondArtistResult.Name, artist.Name)
}

func TestAllArtists(t *testing.T) {
	artist1 := Artist{
		Name: "Test All Artists 1",
	}
	artist2 := Artist{
		Name: "Test All Artists 2",
	}
	artist3 := Artist{
		Name: "Test All Artists 3",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	err = artist1.Create()
	assert.NoError(t, err)
	err = artist2.Create()
	assert.NoError(t, err)
	err = artist3.Create()
	assert.NoError(t, err)

	allArtists, err := artist1.All()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(allArtists))
	assert.Equal(t, artist1.Name, allArtists[0].Name)
	assert.Equal(t, artist1.ID, allArtists[0].ID)
	assert.Equal(t, artist2.Name, allArtists[1].Name)
	assert.Equal(t, artist2.ID, allArtists[1].ID)
	assert.Equal(t, artist3.Name, allArtists[2].Name)
	assert.Equal(t, artist3.ID, allArtists[2].ID)

	err = artist1.Delete()
	assert.NoError(t, err)
	err = artist2.Delete()
	assert.NoError(t, err)
	err = artist3.Delete()
	assert.NoError(t, err)
}

func TestUpsertArtist(t *testing.T) {
	artist := Artist{
		Name: "Test Upsert Artist",
	}

	err := SetupDatabase(t)
	assert.NoError(t, err)

	_, err = artist.Lookup("Test Upsert Artist")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = artist.Upsert()
	assert.NoError(t, err)

	firstArtistResult, err := artist.Lookup(artist.Name)
	assert.NoError(t, err)

	artist.Name = "New Name"
	err = artist.Upsert()
	assert.NoError(t, err)

	secondArtistResult, err := artist.Lookup(artist.Name)
	assert.NoError(t, err)

	_, err = artist.Lookup("Test Upsert Artist")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = artist.Delete()
	assert.NoError(t, err)

	assert.NotEqual(t, firstArtistResult.Name, secondArtistResult.Name)
	assert.NotEqual(t, firstArtistResult.Name, artist.Name)
	assert.Equal(t, secondArtistResult.Name, artist.Name)
}
