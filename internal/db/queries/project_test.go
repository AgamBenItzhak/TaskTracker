package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries/testmockdb"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertProject(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	newProject := models.NewMockProject()

	mock, err = testmockdb.DBMockInsertProject(mock, newProject)
	require.NoError(t, err)

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

	projectsIDs := []int{1, 2}

	mock, err = testmockdb.DBMockGetProjectsIDs(mock, projectsIDs)
	require.NoError(t, err)

	projects, err := GetProjectsIDs(context.Background(), mock)
	require.NoError(t, err)
	require.Equal(t, projectsIDs, projects)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newProject := models.NewMockProject()

	mock, err = testmockdb.DBMockGetProjectByID(mock, newProject)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockUpdateProject(mock, newProject)
	require.NoError(t, err)

	err = UpdateProject(context.Background(), mock, newProject)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProject(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	project := models.NewMockProject()

	mock, err = testmockdb.DBMockDeleteProject(mock, project)
	require.NoError(t, err)

	err = DeleteProject(context.Background(), mock, project.ProjectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
