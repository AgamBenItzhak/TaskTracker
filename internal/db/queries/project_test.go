package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertProject(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newProject := models.NewMockProject()

	mock.ExpectQuery("INSERT INTO projects").
		WithArgs(newProject.ProjectName, newProject.Description, newProject.Status, newProject.StartDate, newProject.EndDate).
		WillReturnRows(pgxmock.NewRows([]string{"project_id"}).AddRow(newProject.ProjectID))

	projectID, err := InsertProject(context.Background(), mock, newProject)
	require.NoError(t, err)
	require.Equal(t, newProject.ProjectID, projectID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectsIDs(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mockRows := pgxmock.NewRows([]string{"project_id"}).
		AddRow(1).
		AddRow(2)

	mock.ExpectQuery("SELECT").
		WillReturnRows(mockRows)

	projectIDs, err := GetProjectsIDs(context.Background(), mock)
	require.NoError(t, err)
	require.Equal(t, []int{1, 2}, projectIDs)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newProject := models.NewMockProject()

	mockRows := pgxmock.NewRows([]string{"project_id", "project_name", "description", "status", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(newProject.ProjectID, newProject.ProjectName, newProject.Description, newProject.Status, newProject.StartDate, newProject.EndDate, newProject.CreatedAt, newProject.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newProject.ProjectID).
		WillReturnRows(mockRows)

	project, err := GetProjectByID(context.Background(), mock, newProject.ProjectID)
	require.NoError(t, err)
	require.Equal(t, newProject, project)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateProject(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newProject := models.NewMockProject()

	mock.ExpectExec("UPDATE projects").
		WithArgs(newProject.ProjectName, newProject.Description, newProject.Status, newProject.StartDate, newProject.EndDate, newProject.ProjectID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateProject(context.Background(), mock, newProject)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProject(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectID := 1

	mock.ExpectExec("DELETE FROM projects").
		WithArgs(projectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProject(context.Background(), mock, projectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
