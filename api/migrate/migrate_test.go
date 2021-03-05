package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupAndMigrateDatabaseDB(t *testing.T) {
	err := SetupAndMigrateDatabase()
	assert.NoError(t, err)
}
