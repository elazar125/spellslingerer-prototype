package main

import (
	"api/db"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrateDB(t *testing.T) {
	db.InitDatabase()
	err := migrateDatabase()
	assert.NoError(t, err)
}
