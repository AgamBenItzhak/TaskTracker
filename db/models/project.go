package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Project struct {
	ProjectID   int       `validate:"required" json:"project_id" mapstructure:"project_id"`
	ProjectName string    `validate:"required" json:"project_name" mapstructure:"project_name"`
	Description string    `validate:"required" json:"description" mapstructure:"description"`
	Status      string    `validate:"required" json:"status" mapstructure:"status"`
	StartDate   time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"start_date" mapstructure:"start_date"`
	EndDate     time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"end_date" mapstructure:"end_date"`
	CreatedAt   time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

type ProjectUser struct {
	UserID    int       `validate:"required" json:"user_id" mapstructure:"user_id"`
	ProjectID int       `validate:"required" json:"project_id" mapstructure:"project_id"`
	Role      string    `validate:"required" json:"role" mapstructure:"role"`
	CreatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

func NewProject(projectID int, projectName, description, status string, startDate, endDate, createdAt, updatedAt time.Time) *Project {
	return &Project{
		ProjectID:   projectID,
		ProjectName: projectName,
		Description: description,
		Status:      status,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewMockProject() *Project {
	return &Project{
		ProjectID:   1,
		ProjectName: "projectName",
		Description: "description",
		Status:      "status",
		StartDate:   time.Now(),
		EndDate:     time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p *Project) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func NewProjectUser(userID, projectID int, role string, createdAt, updatedAt time.Time) *ProjectUser {
	return &ProjectUser{
		UserID:    userID,
		ProjectID: projectID,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewMockProjectUser() *ProjectUser {
	return &ProjectUser{
		UserID:    1,
		ProjectID: 1,
		Role:      "role",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pu *ProjectUser) Validate() error {
	validate := validator.New()
	return validate.Struct(pu)
}
