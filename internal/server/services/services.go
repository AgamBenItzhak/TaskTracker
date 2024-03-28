package services

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/auth"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/project"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/task"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/user"
)

// Services provides functions for the server
type Services struct {
	User    *user.UserService
	Project *project.ProjectService
	Task    *task.TaskService
	Auth    *auth.AuthService
}

// NewServices creates a new Services instance
func NewServices(db db.PgxIface) *Services {
	return &Services{
		User:    user.NewUserService(db),
		Project: project.NewProjectService(db),
		Task:    task.NewTaskService(db),
		Auth:    auth.NewAuthService(db),
	}
}
