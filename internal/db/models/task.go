package models

type TaskGroups struct {
	TaskGroupID int    `json:"task_group_id"`
	ProjectID   int    `json:"project_id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Tasks struct {
	TaskID      int    `json:"task_id"`
	TaskGroupID int    `json:"task_group_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
