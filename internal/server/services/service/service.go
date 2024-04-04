package service

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/pashagolub/pgxmock/v3"
)

// ServiceIface is an interface for the Service
type ServiceIface interface{}

// Service provides service-related functions
type Service struct {
	Queries *db.Queries
}

// Ensure Service implements ServiceIface
var _ ServiceIface = (*Service)(nil)

// NewService creates a new Service instance
func NewService(q db.DBTX) *Service {
	return &Service{Queries: db.New(q)}
}

// NewMockService creates a new Service instance with a mock db connection
func NewMockService() *Service {
	mock, _ := pgxmock.NewPool()
	return NewService(mock)
}
