package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func GenerateToken(secret string, claims ClaimsIface) (string, error) {
	err := claims.Validate()
	if err != nil {
		return "", err
	}

	c := claims.GetClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func VerifyToken(secret string, token string, claims ClaimsIface) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	err = ValidateToken(parsedToken, claims)
	if err != nil {
		return err
	}

	return nil
}

func ValidateToken(token *jwt.Token, claims ClaimsIface) error {
	if !token.Valid {
		return ErrInvalidToken
	}

	err := claims.Validate()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Implement RevokeToken
func RevokeToken(token string) error {
	return nil
}
