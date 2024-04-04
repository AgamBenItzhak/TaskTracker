package project

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/service"
	"github.com/pashagolub/pgxmock/v3"
)

// ServiceIface is an interface for the Service
type ServiceIface interface {
	CreateProject(db.CreateProjectParams) (*db.Project, error)
	GetAllProjects() ([]*db.Project, error)
	GetProjectByID(int32) (*db.Project, error)
	UpdateProjectByID(int32, db.UpdateProjectByIDParams) (*db.Project, error)
	DeleteProjectByID(int32) error

	CreateProjectMember(int32, db.CreateProjectMemberParams) error
	GetAllProjectMembers(int32) ([]*db.ProjectMember, error)
	GetProjectMemberByID(int32, int32) (*db.ProjectMember, error)
	UpdateProjectMemberByID(int32, int32, db.UpdateProjectMemberByIDParams) (*db.ProjectMember, error)
	DeleteProjectMemberByID(int32, int32) error
}

// Service provides project-related functions
type Service struct {
	service *service.Service
}

// Ensure Service implements ServiceIface
var _ ServiceIface = (*Service)(nil)

// NewService creates a new Service instance
func NewService(q db.DBTX) *Service {
	return &Service{service: service.NewService(q)}
}

// NewMockService creates a new Service instance for testing
func NewMockService() *Service {
	mock, _ := pgxmock.NewPool()
	return NewService(mock)
}
