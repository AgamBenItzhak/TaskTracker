package queries

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectID := 1
	name := "name"
	description := "description"

	mock.ExpectQuery("INSERT INTO tasks_groups").
		WithArgs(projectID, name, description).
		WillReturnRows(pgxmock.NewRows([]string{"task_group_id"}).AddRow(1))

	TasksGroupID, err := InsertTasksGroup(context.Background(), mock, projectID, name, description)
	require.NoError(t, err)
	require.Equal(t, 1, TasksGroupID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetTasksGroupsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroupID := 1
	projectID := 1
	name := "name"
	description := "description"
	createdAt := "2021-01-01T00:00:00Z"
	updatedAt := "2021-01-01T00:00:00Z"

	mockRows := pgxmock.NewRows([]string{"project_id", "name", "description", "created_at", "updated_at"}).
		AddRow(projectID, name, description, createdAt, updatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(TasksGroupID).
		WillReturnRows(mockRows)

	pID, n, d, cA, uA, err := GetTasksGroupsByID(context.Background(), mock, TasksGroupID)
	require.NoError(t, err)
	require.Equal(t, projectID, pID)
	require.Equal(t, name, n)
	require.Equal(t, description, d)
	require.Equal(t, createdAt, cA)
	require.Equal(t, updatedAt, uA)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroupID := 1
	name := "name"
	description := "description"

	mock.ExpectExec("UPDATE tasks_groups").
		WithArgs(TasksGroupID, name, description).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateTasksGroup(context.Background(), mock, TasksGroupID, name, description)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroupID := 1

	mock.ExpectExec("DELETE FROM tasks_groups").
		WithArgs(TasksGroupID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteTasksGroup(context.Background(), mock, TasksGroupID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroupID := 1
	name := "name"
	description := "description"
	status := "status"
	priority := "priority"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"

	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(TasksGroupID, name, description, status, priority, startDate, endDate).
		WillReturnRows(pgxmock.NewRows([]string{"task_id"}).AddRow(1))

	taskID, err := InsertTask(context.Background(), mock, TasksGroupID, name, description, status, priority, startDate, endDate)
	require.NoError(t, err)
	require.Equal(t, 1, taskID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskID := 1
	TasksGroupID := 1
	name := "name"
	description := "description"
	status := "status"
	priority := "priority"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"
	createdAt := "2021-01-01T00:00:00Z"
	updatedAt := "2021-01-01T00:00:00Z"

	mockRows := pgxmock.NewRows([]string{"task_group_id", "name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(TasksGroupID, name, description, status, priority, startDate, endDate, createdAt, updatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(taskID).
		WillReturnRows(mockRows)

	tGID, n, d, s, p, sD, eD, uA, err := GetTaskByID(context.Background(), mock, taskID)
	require.NoError(t, err)
	require.Equal(t, TasksGroupID, tGID)
	require.Equal(t, name, n)
	require.Equal(t, description, d)
	require.Equal(t, status, s)
	require.Equal(t, priority, p)
	require.Equal(t, startDate, sD)
	require.Equal(t, endDate, eD)
	require.Equal(t, updatedAt, uA)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskID := 1
	name := "name"
	description := "description"
	status := "status"
	priority := "priority"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"

	mock.ExpectExec("UPDATE tasks").
		WithArgs(taskID, name, description, status, priority, startDate, endDate).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateTaskByID(context.Background(), mock, taskID, name, description, status, priority, startDate, endDate)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskID := 1

	mock.ExpectExec("DELETE FROM tasks").
		WithArgs(taskID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteTaskByID(context.Background(), mock, taskID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
