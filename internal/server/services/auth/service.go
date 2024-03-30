package auth

import (
	"github.com/AgamBenItzhak/TaskTracker/config"
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pashagolub/pgxmock/v3"
)

// AuthService provides auth-related functions
type AuthService struct {
	dbpool      db.PgxIface
	tokenSecret string
}

type AuthServiceIface interface {
	GenerateTokenWithClaims(claims jwt.Claims) (string, error)
	GenerateToken(tokenSecret string, expirationTime int64) (string, error)
	ParseAndVerifyToken(tokenString string) (*Claims, error)
}

type AuthServiceMockIface interface {
	AuthServiceIface
}

type Claims struct {
	jwt.RegisteredClaims
}

// NewAuthService creates a new AuthService instance
func NewAuthService(dbpool db.PgxIface) *AuthService {
	return &AuthService{
		dbpool:      dbpool,
		tokenSecret: config.ServerConfig.JWT.Secret,
	}
}

func NewMockAuthService() *AuthService {
	mock, _ := pgxmock.NewPool()
	return NewAuthService(mock)
}
