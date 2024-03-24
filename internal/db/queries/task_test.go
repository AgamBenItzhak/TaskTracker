package queries

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertTaskGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectID := 1
	name := "name"
	description := "description"

	mock.ExpectQuery("INSERT INTO task_groups").
		WithArgs(projectID, name, description).
		WillReturnRows(pgxmock.NewRows([]string{"task_group_id"}).AddRow(1))

	taskGroupID, err := InsertTaskGroup(context.Background(), mock, projectID, name, description)
	require.NoError(t, err)
	require.Equal(t, 1, taskGroupID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetTaskGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskGroupID := 1
	projectID := 1
	name := "name"
	description := "description"
	createdAt := "2021-01-01T00:00:00Z"
	updatedAt := "2021-01-01T00:00:00Z"

	mockRows := pgxmock.NewRows([]string{"project_id", "name", "description", "created_at", "updated_at"}).
		AddRow(projectID, name, description, createdAt, updatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(taskGroupID).
		WillReturnRows(mockRows)

	pID, n, d, cA, uA, err := GetTaskGroup(context.Background(), mock, taskGroupID)
	require.NoError(t, err)
	require.Equal(t, projectID, pID)
	require.Equal(t, name, n)
	require.Equal(t, description, d)
	require.Equal(t, createdAt, cA)
	require.Equal(t, updatedAt, uA)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateTaskGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskGroupID := 1
	name := "name"
	description := "description"

	mock.ExpectExec("UPDATE task_groups").
		WithArgs(taskGroupID, name, description).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateTaskGroup(context.Background(), mock, taskGroupID, name, description)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTaskGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskGroupID := 1

	mock.ExpectExec("DELETE FROM task_groups").
		WithArgs(taskGroupID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteTaskGroup(context.Background(), mock, taskGroupID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskGroupID := 1
	name := "name"
	description := "description"
	status := "status"
	priority := "priority"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"

	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(taskGroupID, name, description, status, priority, startDate, endDate).
		WillReturnRows(pgxmock.NewRows([]string{"task_id"}).AddRow(1))

	taskID, err := InsertTask(context.Background(), mock, taskGroupID, name, description, status, priority, startDate, endDate)
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
	taskGroupID := 1
	name := "name"
	description := "description"
	status := "status"
	priority := "priority"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"
	createdAt := "2021-01-01T00:00:00Z"
	updatedAt := "2021-01-01T00:00:00Z"

	mockRows := pgxmock.NewRows([]string{"task_group_id", "name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(taskGroupID, name, description, status, priority, startDate, endDate, createdAt, updatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(taskID).
		WillReturnRows(mockRows)

	tGID, n, d, s, p, sD, eD, uA, err := GetTask(context.Background(), mock, taskID)
	require.NoError(t, err)
	require.Equal(t, taskGroupID, tGID)
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

	err = UpdateTask(context.Background(), mock, taskID, name, description, status, priority, startDate, endDate)
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

	err = DeleteTask(context.Background(), mock, taskID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
