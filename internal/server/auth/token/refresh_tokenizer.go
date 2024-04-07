package token

import "github.com/golang-jwt/jwt/v5"

type RefreshableTokenizerIface interface {
	TokenizerIface
	GetRefreshSecret() string
	GenerateTokens(tokenClaims, refreshTokenClaims ClaimsIface) (string, string, error)
	RefreshToken(token string) (string, error)
}

type RefreshableTokenizer struct {
	claims           ClaimsIface
	tokenizer        TokenizerIface
	refreshTokenizer TokenizerIface
}

var _ RefreshableTokenizerIface = (*RefreshableTokenizer)(nil)

func (t *RefreshableTokenizer) GetSecret() string {
	return t.tokenizer.GetSecret()
}

func (t *RefreshableTokenizer) GetRefreshSecret() string {
	return t.refreshTokenizer.GetSecret()
}

func NewRefreshableTokenizer(secret string, refreshSecret string, claims ClaimsIface) *RefreshableTokenizer {
	return &RefreshableTokenizer{
		claims:           claims,
		tokenizer:        NewTokenizer(secret, claims),
		refreshTokenizer: NewTokenizer(refreshSecret, claims),
	}
}

func (t *RefreshableTokenizer) GenerateToken(claims ClaimsIface) (string, error) {
	return t.tokenizer.GenerateToken(claims)
}

func (t *RefreshableTokenizer) VerifyToken(token string) error {
	return t.tokenizer.VerifyToken(token)
}

func (t *RefreshableTokenizer) ValidateToken(token *jwt.Token) error {
	return t.tokenizer.ValidateToken(token)
}

func (t *RefreshableTokenizer) RevokeToken(token string) error {
	return t.tokenizer.RevokeToken(token)
}

func (t *RefreshableTokenizer) GenerateTokens(tokenClaims, refreshTokenClaims ClaimsIface) (string, string, error) {
	token, err := t.tokenizer.GenerateToken(tokenClaims)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := t.refreshTokenizer.GenerateToken(refreshTokenClaims)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (t *RefreshableTokenizer) RefreshToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.refreshTokenizer.GetSecret()), nil
	})
	if err != nil {
		return "", err
	}

	err = t.ValidateToken(parsedToken)
	if err != nil {
		return "", err
	}

	token, err = t.tokenizer.GenerateToken(t.claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
