package testmockdb

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"

	"github.com/pashagolub/pgxmock/v3"
)

func DBMockInsertTasksGroup(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mock.ExpectQuery("INSERT INTO tasks_groups").
		WithArgs(tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description).
		WillReturnRows(pgxmock.NewRows([]string{"tasks_group_id"}).AddRow(tasksGroup.TaskGroupID))

	return mock, nil
}

func DBMockGetAllTasksGroupsIDsByProjectID(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"tasks_group_id"}).
		AddRow(tasksGroup.TaskGroupID)

	mock.ExpectQuery("SELECT").
		WithArgs(tasksGroup.ProjectID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetAllTasksGroupsByProjectID(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"task_group_id", "project_id", "group_name", "description", "created_at", "updated_at"}).
		AddRow(tasksGroup.TaskGroupID, tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description, tasksGroup.CreatedAt, tasksGroup.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(tasksGroup.ProjectID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetTasksGroupsByID(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"task_group_id", "project_id", "group_name", "description", "created_at", "updated_at"}).
		AddRow(tasksGroup.TaskGroupID, tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description, tasksGroup.CreatedAt, tasksGroup.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(tasksGroup.TaskGroupID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockUpdateTasksGroup(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("UPDATE tasks_groups").
		WithArgs(tasksGroup.TaskGroupID, tasksGroup.ProjectID, tasksGroup.GroupName, tasksGroup.Description).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	return mock, nil
}

func DBMockDeleteTasksGroup(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM tasks_groups").
		WithArgs(tasksGroup.TaskGroupID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockDeleteAllTasksGroupsByProjectID(mock pgxmock.PgxPoolIface, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM tasks_groups").
		WithArgs(project.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockInsertTask(mock pgxmock.PgxPoolIface, task *models.Tasks) (pgxmock.PgxPoolIface, error) {
	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate).
		WillReturnRows(pgxmock.NewRows([]string{"task_id"}).AddRow(task.TaskID))

	return mock, nil
}

func DBMockGetTaskByID(mock pgxmock.PgxPoolIface, task *models.Tasks) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"task_id", "task_group_id", "task_name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(task.TaskID, task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate, task.CreatedAt, task.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(task.TaskID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetAllTasksByTasksGroupID(mock pgxmock.PgxPoolIface, task *models.Tasks) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"task_id", "task_group_id", "task_name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(task.TaskID, task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate, task.CreatedAt, task.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(task.TaskGroupID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockUpdateTaskByID(mock pgxmock.PgxPoolIface, task *models.Tasks) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("UPDATE tasks").
		WithArgs(task.TaskID, task.TaskGroupID, task.TaskName, task.Description, task.Status, task.Priority, task.StartDate, task.EndDate).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	return mock, nil
}

func DBMockDeleteAllTasksByTasksGroupID(mock pgxmock.PgxPoolIface, tasksGroup *models.TasksGroups) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM tasks").
		WithArgs(tasksGroup.TaskGroupID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockDeleteTaskByID(mock pgxmock.PgxPoolIface, task *models.Tasks) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM tasks").
		WithArgs(task.TaskID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}
