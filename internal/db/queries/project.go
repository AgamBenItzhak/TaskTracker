package queries

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
)

// InsertProject inserts a new project into the database
func InsertProject(ctx context.Context, dbpool PgxIface, project *models.Project) (int, error) {
	var projectID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO projects (project_name, description, status, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING project_id
	`, project.ProjectName, project.Description, project.Status, project.StartDate, project.EndDate).Scan(&projectID)
	return projectID, err
}

// GetProjects retrieves all projects from the database
func GetProjectsIDs(ctx context.Context, dbpool PgxIface) ([]int, error) {
	var projectIDs []int
	rows, err := dbpool.Query(ctx, `
		SELECT project_id
		FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var projectID int
		err := rows.Scan(&projectID)
		if err != nil {
			return nil, err
		}
		projectIDs = append(projectIDs, projectID)
	}
	return projectIDs, nil
}

// GetProject retrieves a project from the database
func GetProjectByID(ctx context.Context, dbpool PgxIface, projectID int) (*models.Project, error) {
	var project models.Project
	err := dbpool.QueryRow(ctx, `
		SELECT project_id, project_name, description, status, start_date, end_date, created_at, updated_at
		FROM projects
		WHERE project_id = $1
	`, projectID).Scan(&project.ProjectID, &project.ProjectName, &project.Description, &project.Status, &project.StartDate, &project.EndDate, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject updates a project in the database
func UpdateProject(ctx context.Context, dbpool PgxIface, project *models.Project) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE projects
		SET project_name = $1, description = $2, status = $3, start_date = $4, end_date = $5, updated_at = CURRENT_TIMESTAMP
		WHERE project_id = $6
	`, project.ProjectName, project.Description, project.Status, project.StartDate, project.EndDate, project.ProjectID)
	return err
}

// DeleteProject deletes a project from the database
func DeleteProject(ctx context.Context, dbpool PgxIface, projectID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM projects
		WHERE project_id = $1
	`, projectID)
	return err
}
