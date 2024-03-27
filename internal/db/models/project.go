package models

type Project struct {
	ProjectID   int    `validate:"required" json:"project_id" mapstructure:"project_id"`
	ProjectName string `validate:"required" json:"project_name" mapstructure:"project_name"`
	Description string `json:"description" mapstructure:"description"`
	Status      string `validate:"required" json:"status" mapstructure:"status"`
	StartDate   string `validate:"date" json:"start_date" mapstructure:"start_date"`
	EndDate     string `validate:"date" json:"end_date" mapstructure:"end_date"`
	CreatedAt   string `validate:"required,date" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   string `validate:"required,date" json:"updated_at" mapstructure:"updated_at"`
}

func NewProject(projectID int, projectName, description, status, startDate, endDate, createdAt, updatedAt string) *Project {
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
		StartDate:   "2021-01-01T00:00:00Z",
		EndDate:     "2021-01-01T00:00:00Z",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
	}
}
