package queries

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
)

// InsertTasksGroup inserts a new task group into the database
func InsertTasksGroup(ctx context.Context, dbpool db.PgxIface, tasksGroup *models.TasksGroups) (int, error) {
	var tasksGroupID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO tasks_groups (project_id, group_name, description, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_group_id
	`, tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description).Scan(&tasksGroupID)
	return tasksGroupID, err
}

// GetAllTasksGroupsIDsByProjectID retrieves all task groups IDs from the database
func GetAllTasksGroupsIDsByProjectID(ctx context.Context, dbpool db.PgxIface, projectID int) ([]int, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT task_group_id
		FROM tasks_groups
		WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var TasksGroupIDs []int
	for rows.Next() {
		var TasksGroupID int
		err = rows.Scan(&TasksGroupID)
		if err != nil {
			return nil, err
		}
		TasksGroupIDs = append(TasksGroupIDs, TasksGroupID)
	}
	return TasksGroupIDs, nil
}

// GetAllTasksGroupsByProjectID retrieves all task groups from the database
func GetAllTasksGroupsByProjectID(ctx context.Context, dbpool db.PgxIface, projectID int) ([]*models.TasksGroups, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT task_group_id, project_id, group_name, description, created_at, updated_at
		FROM tasks_groups
		WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasksGroups []*models.TasksGroups
	for rows.Next() {
		var taskGroupID, projectID int
		var groupName, description, createdAt, updatedAt string
		err = rows.Scan(&taskGroupID, &projectID, &groupName, &description, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		tasksGroup := models.NewTasksGroups(taskGroupID, projectID, groupName, description, createdAt, updatedAt)

		err = tasksGroup.Validate()
		if err != nil {
			return nil, err
		}

		tasksGroups = append(tasksGroups, tasksGroup)
	}
	return tasksGroups, nil
}

// GetTasksGroupsByID retrieves a task group from the database
func GetTasksGroupsByID(ctx context.Context, dbpool db.PgxIface, TasksGroupID int) (*models.TasksGroups, error) {
	var taskGroupID, projectID int
	var groupName, description, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT task_group_id, project_id, group_name, description, created_at, updated_at
		FROM tasks_groups
		WHERE task_group_id = $1
	`, TasksGroupID).Scan(&taskGroupID, &projectID, &groupName, &description, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	tasksGroup := models.NewTasksGroups(taskGroupID, projectID, groupName, description, createdAt, updatedAt)

	err = tasksGroup.Validate()
	if err != nil {
		return nil, err
	}

	return tasksGroup, nil
}

// UpdateTasksGroup updates a task group in the database
func UpdateTasksGroup(ctx context.Context, dbpool db.PgxIface, tasksGroup *models.TasksGroups) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE tasks_groups
		SET project_id = $2, group_name = $3, description = $4, updated_at = CURRENT_TIMESTAMP
		WHERE task_group_id = $1
	`, tasksGroup.TaskGroupID, tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description)
	return err
}

// DeleteTasksGroup deletes a task group from the database
func DeleteTasksGroup(ctx context.Context, dbpool db.PgxIface, TasksGroupID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks_groups
		WHERE task_group_id = $1
	`, TasksGroupID)
	return err
}

func DeleteAllTasksGroupsByProjectID(ctx context.Context, dbpool db.PgxIface, projectID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks_groups
		WHERE project_id = $1
	`, projectID)
	return err
}

// InsertTask inserts a new task into the database
func InsertTask(ctx context.Context, dbpool db.PgxIface, task *models.Tasks) (int, error) {
	var taskID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO tasks (task_group_id, task_name, description, status, priority, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING task_id
	`, task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate).Scan(&taskID)
	return taskID, err
}

// GetTask retrieves a task from the database
func GetTaskByID(ctx context.Context, dbpool db.PgxIface, taskID int) (*models.Tasks, error) {
	var taskGroupID int
	var taskName, description, status, priority, startDate, endDate, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT task_id, task_group_id, task_name, description, status, priority, start_date, end_date, created_at, updated_at
		FROM tasks
		WHERE task_id = $1
	`, taskID).Scan(&taskID, &taskGroupID, &taskName, &description, &status, &priority, &startDate, &endDate, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	task := models.NewTasks(taskID, taskGroupID, taskName, description, status, priority, startDate, endDate, createdAt, updatedAt)

	err = task.Validate()
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTasksByTasksGroupID retrieves all tasks from the database
func GetAllTasksByTasksGroupID(ctx context.Context, dbpool db.PgxIface, TasksGroupID int) ([]*models.Tasks, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT task_id, task_group_id, task_name, description, status, priority, start_date, end_date, created_at, updated_at
		FROM tasks
		WHERE task_group_id = $1
	`, TasksGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Tasks
	for rows.Next() {
		var taskID, taskGroupID int
		var taskName, description, status, priority, startDate, endDate, createdAt, updatedAt string
		err = rows.Scan(&taskID, &taskGroupID, &taskName, &description, &status, &priority, &startDate, &endDate, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		task := models.NewTasks(taskID, taskGroupID, taskName, description, status, priority, startDate, endDate, createdAt, updatedAt)

		err = task.Validate()
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTaskByID updates a task in the database
func UpdateTaskByID(ctx context.Context, dbpool db.PgxIface, task *models.Tasks) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE tasks
		SET task_group_id = $2, task_name = $3, description = $4, status = $5, priority = $6, start_date = $7, end_date = $8, updated_at = CURRENT_TIMESTAMP
		WHERE task_id = $1
	`, task.TaskID, task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate)
	return err
}

// DeleteAllTasksByTasksGroupID deletes all tasks from the database
func DeleteAllTasksByTasksGroupID(ctx context.Context, dbpool db.PgxIface, TasksGroupID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks
		WHERE task_group_id = $1
	`, TasksGroupID)
	return err
}

// DeleteTask deletes a task from the database
func DeleteTaskByID(ctx context.Context, dbpool db.PgxIface, taskID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM tasks
		WHERE task_id = $1
	`, taskID)
	return err
}
