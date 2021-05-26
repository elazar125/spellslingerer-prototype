package filter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastStringToType(t *testing.T) {
	result, err := castStringToType("123", reflect.ValueOf(0).Type())
	assert.NoError(t, err)
	assert.Equal(t, 123, result)

	result, err = castStringToType("123", reflect.ValueOf("").Type())
	assert.NoError(t, err)
	assert.Equal(t, "123", result)

	type newType interface{}
	result, err = castStringToType("123", reflect.ValueOf(new(newType)).Type())
	assert.EqualError(t, err, "Unrecognized type")
}

type testType struct {
	name   string `filter-alias:"C" filter-column-name:"nombre"`
	number int    `filter-alias:"C" filter-column-name:"numero"`
}

func TestCastQueryParameters(t *testing.T) {
	initial := map[string][]string{
		"name":   {"test"},
		"number": {"1"},
	}
	expected := map[string][]interface{}{
		"C.nombre": {"test"},
		"C.numero": {1},
	}
	result, err := castQueryParameters(initial, testType{})
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestBuildWhereClauseFromQuery(t *testing.T) {
	initial := map[string][]string{
		"name":   {"test"},
		"number": {"1", "2"},
	}
	expectedClause := "WHERE C.nombre IN ($1) AND C.numero IN ($2, $3)"
	expectedArgs := []interface{}{"test", 1, 2}

	clause, args, err := BuildWhereClauseFromQuery(initial, testType{})
	assert.Equal(t, expectedClause, clause)
	assert.Equal(t, expectedArgs, args)
	assert.NoError(t, err)
}

func TestBuildWhereClauseFromQueryEmpty(t *testing.T) {
	initial := map[string][]string{}
	expectedClause := ""
	expectedArgs := []interface{}(nil)

	clause, args, err := BuildWhereClauseFromQuery(initial, testType{})
	assert.Equal(t, expectedClause, clause)
	assert.Equal(t, expectedArgs, args)
	assert.NoError(t, err)
}
