package auth

import (
	"api/db"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJwtTokenFromAuthHeader(t *testing.T) {
	token, err := GetJwtTokenFromAuthHeader("Invalid Authorization Header")
	assert.Nil(t, token)
	assert.Error(t, err)

	token, err = GetJwtTokenFromAuthHeader("Bearer <token>")
	assert.Equal(t, "<token>", token.Token)
	assert.NoError(t, err)
}

func TestGenerateJwtToken(t *testing.T) {
	os.Setenv("JWT_ISSUER", "TestGenerateJwtToken")
	os.Setenv("JWT_KEY", "TestGenerateJwtToken")

	generatedToken, err := GenerateJwtToken("TestValidateJwtToken@email.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, generatedToken.Token)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}

func TestBlacklist(t *testing.T) {
	os.Setenv("JWT_ISSUER", "TestBlacklist")
	os.Setenv("JWT_KEY", "TestBlacklist")

	generatedToken, err := GenerateJwtToken("TestValidateJwtToken@email.com")
	assert.NoError(t, err)
	assert.NotNil(t, generatedToken)

	err = db.InitDatabase()
	assert.NoError(t, err)

	err = generatedToken.Blacklist()
	assert.NoError(t, err)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}

func TestIsBlacklisted(t *testing.T) {
	os.Setenv("JWT_ISSUER", "TestIsBlacklisted")
	os.Setenv("JWT_KEY", "TestIsBlacklisted")

	generatedToken, err := GenerateJwtToken("TestValidateJwtToken@email.com")
	assert.NoError(t, err)
	assert.NotNil(t, generatedToken)

	err = db.InitDatabase()
	assert.NoError(t, err)

	isBlacklisted, err := generatedToken.isBlacklisted()
	assert.NoError(t, err)
	assert.False(t, isBlacklisted)

	err = generatedToken.Blacklist()
	assert.NoError(t, err)

	isBlacklisted, err = generatedToken.isBlacklisted()
	assert.NoError(t, err)
	assert.True(t, isBlacklisted)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}

func TestParse(t *testing.T) {
	os.Setenv("JWT_ISSUER", "TestParse")
	os.Setenv("JWT_KEY", "TestParse")

	generatedToken, err := GenerateJwtToken("TestValidateJwtToken@email.com")
	assert.NoError(t, err)
	assert.NotNil(t, generatedToken)

	_, err = generatedToken.parse()
	assert.NoError(t, err)

	malformedToken := &SignedToken{
		Token: "malformed data",
	}

	_, err = malformedToken.parse()
	assert.Error(t, err)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}

func TestValidateJwtToken(t *testing.T) {
	os.Setenv("JWT_ISSUER", "TestValidateJwtToken")
	os.Setenv("JWT_KEY", "TestValidateJwtToken")

	err := db.InitDatabase()
	assert.NoError(t, err)

	generatedToken, err := GenerateJwtToken("TestValidateJwtToken@email.com")
	assert.NoError(t, err)

	claims, err := generatedToken.validate()
	assert.NoError(t, err)

	assert.Equal(t, "TestValidateJwtToken@email.com", claims.Email)
	assert.Equal(t, "TestValidateJwtToken", claims.Issuer)

	err = generatedToken.Blacklist()
	assert.NoError(t, err)

	_, err = generatedToken.validate()
	assert.Error(t, err)

	malformedToken := &SignedToken{
		Token: "malformed data",
	}

	_, err = malformedToken.parse()
	assert.Error(t, err)

	os.Unsetenv("JWT_ISSUER")
	os.Unsetenv("JWT_KEY")
}
