package queries

import (
	"context"
)

// InsertTaskGroup inserts a new task group into the database
func InsertTaskGroup(ctx context.Context, dbpool PgxIface, projectID int, name string, description string) (int, error) {
	var taskGroupID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO task_groups (project_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_group_id
	`, projectID, name, description).Scan(&taskGroupID)
	return taskGroupID, err
}

// GetTaskGroup retrieves a task group from the database
func GetTaskGroup(ctx context.Context, dbpool PgxIface, taskGroupID int) (int, string, string, string, string, error) {
	var projectID int
	var name, description, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT project_id, name, description, created_at, updated_at
		FROM task_groups
		WHERE task_group_id = $1
	`, taskGroupID).Scan(&projectID, &name, &description, &createdAt, &updatedAt)
	return projectID, name, description, createdAt, updatedAt, err
}

// UpdateTaskGroup updates a task group in the database
func UpdateTaskGroup(ctx context.Context, dbpool PgxIface, taskGroupID int, name string, description string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE task_groups
		SET name = $2, description = $3, updated_at = CURRENT_TIMESTAMP
		WHERE task_group_id = $1
	`, taskGroupID, name, description)
	return err
}

// DeleteTaskGroup deletes a task group from the database
func DeleteTaskGroup(ctx context.Context, dbpool PgxIface, taskGroupID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM task_groups
		WHERE task_group_id = $1
	`, taskGroupID)
	return err
}

// InsertTask inserts a new task into the database
func InsertTask(ctx context.Context, dbpool PgxIface, taskGroupID int, name string, description string, status string, priority string, startDate string, endDate string) (int, error) {
	var taskID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO tasks (task_group_id, name, description, status, priority, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_id
	`, taskGroupID, name, description, status, priority, startDate, endDate).Scan(&taskID)
	return taskID, err
}

// GetTask retrieves a task from the database
func GetTask(ctx context.Context, dbpool PgxIface, taskID int) (int, string, string, string, string, string, string, string, error) {
	var taskGroupID int
	var name, description, status, priority, startDate, endDate, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT task_group_id, name, description, status, priority, start_date, end_date, created_at, updated_at
		FROM tasks
		WHERE task_id = $1
	`, taskID).Scan(&taskGroupID, &name, &description, &status, &priority, &startDate, &endDate, &createdAt, &updatedAt)
	return taskGroupID, name, description, status, priority, startDate, endDate, updatedAt, err
}

// UpdateTask updates a task in the database
func UpdateTask(ctx context.Context, dbpool PgxIface, taskID int, name string, description string, status string, priority string, startDate string, endDate string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE tasks
		SET name = $2, description = $3, status = $4, priority = $5, start_date = $6, end_date = $7, updated_at = CURRENT_TIMESTAMP
		WHERE task_id = $1
	`, taskID, name, description, status, priority, startDate, endDate)
	return err
}

// DeleteTask deletes a task from the database
func DeleteTask(ctx context.Context, dbpool PgxIface, taskID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks
		WHERE task_id = $1
	`, taskID)
	return err
}
