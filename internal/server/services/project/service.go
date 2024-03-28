package project

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/pashagolub/pgxmock/v3"
)

// ProjectService provides project-related functions
type ProjectService struct {
	dbpool db.PgxIface
}

// NewProjectService creates a new ProjectService instance
func NewProjectService(dbpool db.PgxIface) *ProjectService {
	return &ProjectService{dbpool: dbpool}
}

func NewMockProjectService() *ProjectService {
	mock, _ := pgxmock.NewPool()
	return NewProjectService(mock)
}
