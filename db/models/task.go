package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TasksGroup struct {
	TasksGroupID int       `validate:"required" json:"task_group_id" mapstructure:"task_group_id"`
	ProjectID    int       `validate:"required" json:"project_id" mapstructure:"project_id"`
	GroupName    string    `validate:"required" json:"group_name" mapstructure:"group_name"`
	Description  string    `validate:"required" json:"description" mapstructure:"description"`
	Status       string    `validate:"required" json:"status" mapstructure:"status"`
	Priority     string    `validate:"required" json:"priority" mapstructure:"priority"`
	StartDate    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"start_date" mapstructure:"start_date"`
	EndDate      time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"end_date" mapstructure:"end_date"`
	CreatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

type Task struct {
	TaskID       int       `validate:"required" json:"task_id" mapstructure:"task_id"`
	TasksGroupID int       `validate:"required" json:"task_group_id" mapstructure:"task_group_id"`
	TaskName     string    `validate:"required" json:"task_name" mapstructure:"task_name"`
	Description  string    `validate:"required" json:"description" mapstructure:"description"`
	Status       string    `validate:"required" json:"status" mapstructure:"status"`
	Priority     string    `validate:"required" json:"priority" mapstructure:"priority"`
	StartDate    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"start_date" mapstructure:"start_date"`
	EndDate      time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"end_date" mapstructure:"end_date"`
	CreatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

type TasksGroupUser struct {
	UserID       int       `validate:"required" json:"user_id" mapstructure:"user_id"`
	TasksGroupID int       `validate:"required" json:"task_group_id" mapstructure:"task_group_id"`
	Role         string    `validate:"required" json:"role" mapstructure:"role"`
	CreatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

type TaskUser struct {
	UserID    int       `validate:"required" json:"user_id" mapstructure:"user_id"`
	TaskID    int       `validate:"required" json:"task_id" mapstructure:"task_id"`
	Role      string    `validate:"required" json:"role" mapstructure:"role"`
	CreatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

func NewTasksGroup(taskGroupID, projectID int, groupName, description string, createdAt, updatedAt time.Time) *TasksGroup {
	return &TasksGroup{
		TasksGroupID: taskGroupID,
		ProjectID:    projectID,
		GroupName:    groupName,
		Description:  description,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func NewMockTasksGroup() *TasksGroup {
	return &TasksGroup{
		TasksGroupID: 1,
		ProjectID:    1,
		GroupName:    "groupName",
		Description:  "description",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (t *TasksGroup) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func NewTasks(taskID, taskGroupID int, taskName, description, status, priority string, startDate, endDate time.Time) *Task {
	return &Task{
		TaskID:       taskID,
		TasksGroupID: taskGroupID,
		TaskName:     taskName,
		Description:  description,
		Status:       status,
		Priority:     priority,
		StartDate:    startDate,
		EndDate:      endDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewMockTasks() *Task {
	return &Task{
		TaskID:       1,
		TasksGroupID: 1,
		TaskName:     "taskName",
		Description:  "description",
		Status:       "status",
		Priority:     "priority",
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func NewTasksGroupUser(userID, taskGroupID int, role string, createdAt, updatedAt time.Time) *TasksGroupUser {
	return &TasksGroupUser{
		UserID:       userID,
		TasksGroupID: taskGroupID,
		Role:         role,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func NewMockTasksGroupUser() *TasksGroupUser {
	return &TasksGroupUser{
		UserID:       1,
		TasksGroupID: 1,
		Role:         "role",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (t *TasksGroupUser) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func NewTaskUser(userID, taskID int, role string, createdAt, updatedAt time.Time) *TaskUser {
	return &TaskUser{
		UserID:    userID,
		TaskID:    taskID,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewMockTaskUser() *TaskUser {
	return &TaskUser{
		UserID:    1,
		TaskID:    1,
		Role:      "role",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *TaskUser) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
