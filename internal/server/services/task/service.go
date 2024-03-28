package task

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/pashagolub/pgxmock/v3"
)

// TaskService provides task-related functions
type TaskService struct {
	dbpool db.PgxIface
}

// NewTaskService creates a new TaskService instance
func NewTaskService(dbpool db.PgxIface) *TaskService {
	return &TaskService{dbpool: dbpool}
}

func NewMockTaskService() *TaskService {
	mock, _ := pgxmock.NewPool()
	return NewTaskService(mock)
}
