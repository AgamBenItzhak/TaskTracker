package member

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/service"
	"github.com/pashagolub/pgxmock/v3"
)

// ServiceIface is an interface for the Service
type ServiceIface interface {
	CreateMember(db.CreateMemberParams) (*db.Member, error)
	GetAllMembers() ([]*db.Member, error)
	GetMemberByID(int32) (*db.Member, error)
	UpdateMemberByID(int32, db.UpdateMemberByIDParams) (*db.Member, error)
	DeleteMemberByID(int32) error

	CreateMemberCredentials(db.CreateMemberCredentialsParams) error
	GetMemberCredentialsByID(int32) (*db.MemberCredentials, error)
	UpdateMemberCredentialsByID(int32, db.UpdateMemberCredentialsByIDParams) error
	DeleteMemberCredentialsByID(int32) error
}

// Service provides member-related functions
type Service struct {
	service *service.Service
}

// NewService creates a new Service instance
func NewService(q db.DBTX) *Service {
	return &Service{service: service.NewService(q)}
}

// Ensure Service implements ServiceIface
var _ ServiceIface = (*Service)(nil)

// NewMockService creates a new Service instance for testing
func NewMockService() *Service {
	mock, _ := pgxmock.NewPool()
	return NewService(mock)
}
