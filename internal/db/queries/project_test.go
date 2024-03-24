package queries

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertProject(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	name := "name"
	description := "description"
	status := "status"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"

	mock.ExpectQuery("INSERT INTO projects").
		WithArgs(name, description, status, startDate, endDate).
		WillReturnRows(pgxmock.NewRows([]string{"project_id"}).AddRow(1))

	projectID, err := InsertProject(context.Background(), mock, name, description, status, startDate, endDate)
	require.NoError(t, err)
	require.Equal(t, 1, projectID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProject(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	projectID := 1
	name := "name"
	description := "description"
	status := "status"
	createdAt := "2021-01-01T00:00:00Z"
	updatedAt := "2021-01-01T00:00:00Z"

	mockRows := pgxmock.NewRows([]string{"name", "description", "status", "created_at", "updated_at"}).
		AddRow(name, description, status, createdAt, updatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(projectID).
		WillReturnRows(mockRows)

	n, d, s, cA, uA, err := GetProject(context.Background(), mock, projectID)
	require.NoError(t, err)
	require.Equal(t, name, n)
	require.Equal(t, description, d)
	require.Equal(t, status, s)
	require.Equal(t, createdAt, cA)
	require.Equal(t, updatedAt, uA)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateProject(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	projectID := 1
	name := "name"
	description := "description"
	status := "status"
	startDate := "2021-01-01T00:00:00Z"
	endDate := "2021-01-01T00:00:00Z"

	mock.ExpectExec("UPDATE projects").
		WithArgs(projectID, name, description, status, startDate, endDate).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateProject(context.Background(), mock, projectID, name, description, status, startDate, endDate)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProject(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())

	projectID := 1

	mock.ExpectExec("DELETE FROM projects").
		WithArgs(projectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProject(context.Background(), mock, projectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
