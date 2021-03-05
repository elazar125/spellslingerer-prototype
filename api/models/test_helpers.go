package models

import (
	"api/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SetupDatabase sets up the database for use in unit tests
// TODO: mock the database
func SetupDatabase(t *testing.T) error {
	t.Log("setup database")

	err := db.InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = AutoMigrateModels()
	assert.NoError(t, err)

	return err
}
