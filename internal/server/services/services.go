package services

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/auth"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/member"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/project"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/task"
)

// Services provides functions for the server
type Services struct {
	Member  *member.Service
	Project *project.Service
	Task    *task.Service
	Auth    *auth.Service
}

// NewServices creates a new Services instance
func NewServices(db db.DBTX, token string) *Services {
	return &Services{
		Member:  member.NewService(db),
		Project: project.NewService(db),
		Task:    task.NewService(db),
		Auth:    auth.NewService(db, token),
	}
}
