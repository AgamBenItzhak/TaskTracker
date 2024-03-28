package testmockdb

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"

	"github.com/pashagolub/pgxmock/v3"
)

func DBMockInsertProject(mock pgxmock.PgxPoolIface, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mock.ExpectQuery("INSERT INTO projects").
		WithArgs(project.ProjectName, project.Description, project.Status, project.StartDate, project.EndDate).
		WillReturnRows(pgxmock.NewRows([]string{"project_id"}).AddRow(project.ProjectID))

	return mock, nil
}

func DBMockGetProjectsIDs(mock pgxmock.PgxPoolIface, projectsIDs []int) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"project_id"})

	for _, project := range projectsIDs {
		mockRows.AddRow(project)
	}

	mock.ExpectQuery("SELECT project_id FROM projects").
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetProjectByID(mock pgxmock.PgxPoolIface, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"project_id", "project_name", "description", "status", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(project.ProjectID, project.ProjectName, project.Description, project.Status, project.StartDate, project.EndDate, project.CreatedAt, project.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(project.ProjectID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockUpdateProject(mock pgxmock.PgxPoolIface, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("UPDATE projects").
		WithArgs(project.ProjectName, project.Description, project.Status, project.StartDate, project.EndDate, project.ProjectID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	return mock, nil
}

func DBMockDeleteProject(mock pgxmock.PgxPoolIface, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM projects").
		WithArgs(project.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}
