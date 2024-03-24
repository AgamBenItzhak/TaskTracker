package queries

import (
	"context"
)

// InsertProject inserts a new project into the database
func InsertProject(ctx context.Context, dbpool PgxIface, name, description string, status string, startDate string, endDate string) (int, error) {
	var projectID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO projects (name, description, status, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING project_id
	`, name, description, status, startDate, endDate).Scan(&projectID)
	return projectID, err
}

// GetProject retrieves a project from the database
func GetProject(ctx context.Context, dbpool PgxIface, projectID int) (string, string, string, string, string, error) {
	var name, description, status, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT name, description, status, created_at, updated_at
		FROM projects
		WHERE project_id = $1
	`, projectID).Scan(&name, &description, &status, &createdAt, &updatedAt)
	return name, description, status, createdAt, updatedAt, err
}

// UpdateProject updates a project in the database
func UpdateProject(ctx context.Context, dbpool PgxIface, projectID int, name, description string, status string, startDate string, endDate string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE projects
		SET name = $2, description = $3, status = $4, start_date = $5, end_date = $6, updated_at = CURRENT_TIMESTAMP
		WHERE project_id = $1
	`, projectID, name, description, status, startDate, endDate)
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
