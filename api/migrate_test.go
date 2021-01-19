package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupAndMigrateDatabaseDB(t *testing.T) {
	err := setupAndMigrateDatabase()
	assert.NoError(t, err)
}
