package models

import (
	"github.com/go-playground/validator/v10"
)

type TasksGroups struct {
	TaskGroupID int    `validate:"required" json:"task_group_id" mapstructure:"task_group_id"`
	ProjectID   int    `validate:"required" json:"project_id" mapstructure:"project_id"`
	GroupName   string `validate:"required" json:"group_name" mapstructure:"group_name"`
	Description string `json:"description" mapstructure:"description"`
	CreatedAt   string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

type Tasks struct {
	TaskID      int    `validate:"required" json:"task_id" mapstructure:"task_id"`
	TaskGroupID int    `validate:"required" json:"task_group_id" mapstructure:"task_group_id"`
	TaskName    string `validate:"required" json:"task_name" mapstructure:"task_name"`
	Description string `json:"description" mapstructure:"description"`
	Status      string `validate:"required" json:"status" mapstructure:"status"`
	Priority    string `validate:"required" json:"priority" mapstructure:"priority"`
	StartDate   string `validate:"datetime=2006-01-02T15:04:05Z07:00" json:"start_date" mapstructure:"start_date"`
	EndDate     string `validate:"datetime=2006-01-02T15:04:05Z07:00" json:"end_date" mapstructure:"end_date"`
	CreatedAt   string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

func NewTasksGroups(taskGroupID, projectID int, groupName, description, createdAt, updatedAt string) *TasksGroups {
	return &TasksGroups{
		TaskGroupID: taskGroupID,
		ProjectID:   projectID,
		GroupName:   groupName,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewMockTasksGroups() *TasksGroups {
	return &TasksGroups{
		TaskGroupID: 1,
		ProjectID:   1,
		GroupName:   "groupName",
		Description: "description",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
	}
}

func (t *TasksGroups) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func NewTasks(taskID, taskGroupID int, taskName, description, status, priority, startDate, endDate, createdAt, updatedAt string) *Tasks {
	return &Tasks{
		TaskID:      taskID,
		TaskGroupID: taskGroupID,
		TaskName:    taskName,
		Description: description,
		Status:      status,
		Priority:    priority,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewMockTasks() *Tasks {
	return &Tasks{
		TaskID:      1,
		TaskGroupID: 1,
		TaskName:    "taskName",
		Description: "description",
		Status:      "status",
		Priority:    "priority",
		StartDate:   "2021-01-01T00:00:00Z",
		EndDate:     "2021-01-01T00:00:00Z",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
	}
}

func (t *Tasks) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
