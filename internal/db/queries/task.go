package queries

import (
	"context"
)

// InsertTasksGroup inserts a new task group into the database
func InsertTasksGroup(ctx context.Context, dbpool PgxIface, projectID int, name string, description string) (int, error) {
	var TasksGroupID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO tasks_groups (project_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_group_id
	`, projectID, name, description).Scan(&TasksGroupID)
	return TasksGroupID, err
}

// GetAllTasksGroupsByProjectID retrieves all task groups from the database
func GetAllTasksGroupsByProjectID(ctx context.Context, dbpool PgxIface, projectID int) ([]int, []string, []string, []string, []string, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT task_group_id, name, description, created_at, updated_at
		FROM tasks_groups
		WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	var TasksGroupIDs []int
	var names, descriptions, createdAts, updatedAts []string
	for rows.Next() {
		var TasksGroupID int
		var name, description, createdAt, updatedAt string
		err = rows.Scan(&TasksGroupID, &name, &description, &createdAt, &updatedAt)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		TasksGroupIDs = append(TasksGroupIDs, TasksGroupID)
		names = append(names, name)
		descriptions = append(descriptions, description)
		createdAts = append(createdAts, createdAt)
		updatedAts = append(updatedAts, updatedAt)
	}
	return TasksGroupIDs, names, descriptions, createdAts, updatedAts, nil
}

// GetTasksGroupsByID retrieves a task group from the database
func GetTasksGroupsByID(ctx context.Context, dbpool PgxIface, TasksGroupID int) (int, string, string, string, string, error) {
	var projectID int
	var groupName, description, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT project_id, name, description, created_at, updated_at
		FROM tasks_groups
		WHERE task_group_id = $1
	`, TasksGroupID).Scan(&projectID, &groupName, &description, &createdAt, &updatedAt)
	return projectID, groupName, description, createdAt, updatedAt, err
}

// UpdateTasksGroup updates a task group in the database
func UpdateTasksGroup(ctx context.Context, dbpool PgxIface, TasksGroupID int, name string, description string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE tasks_groups
		SET name = $2, description = $3, updated_at = CURRENT_TIMESTAMP
		WHERE task_group_id = $1
	`, TasksGroupID, name, description)
	return err
}

// DeleteTasksGroup deletes a task group from the database
func DeleteTasksGroup(ctx context.Context, dbpool PgxIface, TasksGroupID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks_groups
		WHERE task_group_id = $1
	`, TasksGroupID)
	return err
}

// InsertTask inserts a new task into the database
func InsertTask(ctx context.Context, dbpool PgxIface, TasksGroupID int, name string, description string, status string, priority string, startDate string, endDate string) (int, error) {
	var taskID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO tasks (task_group_id, name, description, status, priority, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_id
	`, TasksGroupID, name, description, status, priority, startDate, endDate).Scan(&taskID)
	return taskID, err
}

// GetTask retrieves a task from the database
func GetTask(ctx context.Context, dbpool PgxIface, taskID int) (int, string, string, string, string, string, string, string, error) {
	var TasksGroupID int
	var name, description, status, priority, startDate, endDate, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT task_group_id, name, description, status, priority, start_date, end_date, created_at, updated_at
		FROM tasks
		WHERE task_id = $1
	`, taskID).Scan(&TasksGroupID, &name, &description, &status, &priority, &startDate, &endDate, &createdAt, &updatedAt)
	return TasksGroupID, name, description, status, priority, startDate, endDate, updatedAt, err
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
