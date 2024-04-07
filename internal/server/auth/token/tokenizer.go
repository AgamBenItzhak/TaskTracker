package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenizerIface interface {
	GetSecret() string
	GenerateToken(claims ClaimsIface) (string, error)
	VerifyToken(token string) error
	ValidateToken(token *jwt.Token) error
	RevokeToken(token string) error
}

type Tokenizer struct {
	claims ClaimsIface
	secret string
}

var _ TokenizerIface = (*Tokenizer)(nil)

func NewTokenizer(secret string, claims ClaimsIface) *Tokenizer {
	return &Tokenizer{
		claims: claims,
		secret: secret,
	}
}

func (t *Tokenizer) GetSecret() string {
	return t.secret
}

func (t *Tokenizer) GenerateToken(claims ClaimsIface) (string, error) {
	return GenerateToken(t.secret, claims)
}

func (t *Tokenizer) VerifyToken(token string) error {
	return VerifyToken(t.secret, token, t.claims)
}

func (t *Tokenizer) ValidateToken(token *jwt.Token) error {
	return ValidateToken(token, t.claims)
}

func (t *Tokenizer) RevokeToken(token string) error {
	return RevokeToken(token)
}
