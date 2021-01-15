package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJwtToken(t *testing.T) {
	os.Setenv("JWT_ISSUER", "testdata")
	os.Setenv("JWT_KEY", "testdata")

	generatedToken, err := GenerateJwtToken("jwt@email.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, generatedToken)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}

func TestValidateJwtToken(t *testing.T) {
	os.Setenv("JWT_ISSUER", "testdata")
	os.Setenv("JWT_KEY", "testdata")

	generatedToken, err := GenerateJwtToken("jwt@email.com")
	assert.NoError(t, err)
	claims, err := ValidateJwtToken(generatedToken)
	assert.NoError(t, err)

	assert.Equal(t, "jwt@email.com", claims.Email)
	assert.Equal(t, "testdata", claims.Issuer)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}
