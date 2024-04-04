package auth

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/service"
	"github.com/pashagolub/pgxmock/v3"
)

// ServiceIface is an interface for the Service
type ServiceIface interface {
	LoginMember(email, password string) (string, error)
	LogoutMember(token string) error
	RefreshMemberToken(token string) (string, error)
	UpdateMemberCredentialsByID(email, password string) error
}

// Service provides auth-related functions
type Service struct {
	service *service.Service
	token   string
}

// Ensure Service implements ServiceIface
var _ ServiceIface = (*Service)(nil)

// NewService creates a new Service instance
func NewService(q db.DBTX, token string) *Service {
	return &Service{service: service.NewService(q), token: token}
}

// NewMockService creates a new Service instance for testing
func NewMockService(token string) *Service {
	mock, _ := pgxmock.NewPool()
	return NewService(mock, token)
}
