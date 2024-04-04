package task

import (
	"github.com/AgamBenItzhak/TaskTracker/database/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services/service"
	"github.com/pashagolub/pgxmock/v3"
)

// ServiceIface is an interface for the Service
type ServiceIface interface {
	CreateTaskGroup(int32, db.CreateTaskGroupParams) (*db.TaskGroup, error)
	GetAllTaskGroups(int32) ([]*db.TaskGroup, error)
	GetTaskGroupByID(int32, int32) (*db.TaskGroup, error)
	UpdateTaskGroupByID(int32, int32, db.UpdateTaskGroupByIDParams) (*db.TaskGroup, error)
	DeleteTaskGroupByID(int32) error

	CreateTaskGroupMember(int32, int32, db.CreateTaskGroupMemberParams) (*db.TaskGroupMember, error)
	GetAllTaskGroupMembers(int32, int32) ([]*db.TaskGroupMember, error)
	GetTaskGroupMemberByID(int32, int32, int32) (*db.TaskGroupMember, error)
	UpdateTaskGroupMemberByID(int32, int32, int32, db.UpdateTaskGroupMemberByIDParams) (*db.TaskGroupMember, error)
	DeleteTaskGroupMemberByID(int32, int32, int32) error

	CreateTask(int32, int32, db.CreateTaskParams) (*db.Task, error)
	GetAllTasks(int32, int32) ([]*db.Task, error)
	GetTaskByID(int32, int32, int32) (*db.Task, error)
	UpdateTaskByID(int32, int32, int32, db.UpdateTaskByIDParams) (*db.Task, error)
	DeleteTaskByID(int32, int32, int32) error

	CreateTaskMember(int32, int32, int32, db.CreateTaskMemberParams) (*db.TaskMember, error)
	GetAllTaskMembers(int32, int32, int32) ([]*db.TaskMember, error)
	GetTaskMemberByID(int32, int32, int32, int32) (*db.TaskMember, error)
	UpdateTaskMemberByID(int32, int32, int32, int32, db.UpdateTaskMemberByIDParams) (*db.TaskMember, error)
	DeleteTaskMemberByID(int32, int32, int32, int32) error
}

// Service provides task-related functions
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
