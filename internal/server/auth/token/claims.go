package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	claims jwt.MapClaims
}

type ClaimsIface interface {
	Validate() error
	Add(key string, value interface{})
	AddClaims(m map[string]interface{})
	Get(key string) interface{}
	GetClaims() jwt.MapClaims
}

func (c *Claims) Validate() error {
	return nil
}

func (c *Claims) Add(key string, value interface{}) {
	c.claims[key] = value
}

func (c *Claims) AddClaims(m map[string]interface{}) {
	for k, v := range m {
		c.claims[k] = v
	}
}

func (c *Claims) Get(key string) interface{} {
	return c.claims[key]
}

func (c *Claims) GetClaims() jwt.MapClaims {
	return c.claims
}

func InitializeClaims(ttl time.Duration) Claims {
	return Claims{
		claims: jwt.MapClaims{
			"iat": time.Now().Unix(),
			"nbf": time.Now().Unix(),
			"exp": time.Now().Add(ttl).Unix(),
		},
	}
}
