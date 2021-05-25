package auth

import (
	"api/db"
	"time"

	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtClaim struct {
	Email string
	jwt.StandardClaims
}

// SignedToken represents a signed JWT token
type SignedToken struct {
	Token string
}

// GetJwtTokenFromAuthHeader takes the authorization header (in the form "Bearer <token>") and extracts the token
func GetJwtTokenFromAuthHeader(authorizationHeader string) (token *SignedToken, err error) {
	const PREFIX = "Bearer "

	if !strings.HasPrefix(authorizationHeader, PREFIX) {
		err = errors.New("Invalid token format")
		return
	}

	token = &SignedToken{
		Token: strings.TrimPrefix(authorizationHeader, PREFIX),
	}

	return
}

// GenerateJwtToken generates a jwt token
func GenerateJwtToken(email string) (SignedToken, error) {
	claims := &jwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   os.Getenv("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return SignedToken{Token: signedToken}, err
}

// Blacklist stores a copy of the token in the database to keep a list of invalid (logged out) tokens
func (signedToken *SignedToken) Blacklist() error {
	_, err := db.GlobalSqlxDB.Exec("INSERT INTO public.signed_tokens (token) VALUES ($1)", signedToken.Token)
	return err
}

func (signedToken *SignedToken) validate() (*jwtClaim, error) {
	if isBlacklisted, err := signedToken.isBlacklisted(); err != nil {
		return nil, err
	} else if isBlacklisted {
		return nil, errors.New("Token is blacklisted")
	}

	if token, err := signedToken.parse(); err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*jwtClaim); !ok {
		return nil, errors.New("Couldn't parse claims")
	} else {
		return claims, nil
	}
}

func (signedToken *SignedToken) isBlacklisted() (bool, error) {
	result, err := db.GlobalSqlxDB.Query("SELECT * FROM public.signed_tokens ST WHERE ST.token = $1", signedToken.Token)
	return result.Next(), err
}

func (signedToken *SignedToken) parse() (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		signedToken.Token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)
}
