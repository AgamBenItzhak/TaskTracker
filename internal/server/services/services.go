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
	user    *user.UserService
	project *project.ProjectService
	task    *task.TaskService
	auth    *auth.AuthService
}

// NewServices creates a new Services instance
func NewServices(db *db.DB) *Services {
	return &Services{
		user:    user.NewUserService(db),
		project: project.NewProjectService(db),
		task:    task.NewTaskService(db),
		auth:    auth.NewAuthService(db),
	}
}
