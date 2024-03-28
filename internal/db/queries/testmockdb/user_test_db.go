package testmockdb

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"

	"github.com/pashagolub/pgxmock/v3"
)

func DBMockInsertUser(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mock.ExpectQuery("INSERT INTO users").
		WithArgs(user.Email, user.PasswordHash, user.PasswordSalt, user.FirstName, user.LastName).
		WillReturnRows(pgxmock.NewRows([]string{"user_id"}).AddRow(user.UserID))

	return mock, nil
}

func DBMockGetUsers(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id", "email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at", "last_seen"}).
		AddRow(user.UserID, user.Email, user.PasswordHash, user.PasswordSalt, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt, user.LastSeen)

	mock.ExpectQuery("SELECT").
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetUsersIDs(mock pgxmock.PgxPoolIface, user []*models.User) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id"})

	for _, u := range user {
		mockRows.AddRow(u.UserID)
	}

	mock.ExpectQuery("SELECT").
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetUserByID(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id", "email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at", "last_seen"}).
		AddRow(user.UserID, user.Email, user.PasswordHash, user.PasswordSalt, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt, user.LastSeen)

	mock.ExpectQuery("SELECT").
		WithArgs(user.UserID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetUserByEmail(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id", "email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at", "last_seen"}).
		AddRow(user.UserID, user.Email, user.PasswordHash, user.PasswordSalt, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt, user.LastSeen)

	mock.ExpectQuery("SELECT").
		WithArgs(user.Email).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockUpdateUser(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("UPDATE users").
		WithArgs(user.UserID, user.Email, user.PasswordHash, user.PasswordSalt, user.FirstName, user.LastName).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	return mock, nil
}

func DBMockDeleteUserByID(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM users").
		WithArgs(user.UserID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockDeleteUserByEmail(mock pgxmock.PgxPoolIface, user *models.User) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM users").
		WithArgs(user.Email).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockInsertProjectUser(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("INSERT INTO project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID, projectUser.Role).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	return mock, nil
}

func DBMockGetProjectUser(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id", "project_id", "role", "created_at", "updated_at"}).
		AddRow(projectUser.UserID, projectUser.ProjectID, projectUser.Role, projectUser.CreatedAt, projectUser.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(projectUser.UserID, projectUser.ProjectID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockUpdateProjectUser(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("UPDATE project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID, projectUser.Role).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	return mock, nil
}

func DBMockDeleteProjectUser(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockGetProjectsByUserID(mock pgxmock.PgxPoolIface, user *models.User, project *models.Project) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"project_id"}).
		AddRow(project.ProjectID)

	mock.ExpectQuery("SELECT").
		WithArgs(user.UserID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockGetUsersByProjectID(mock pgxmock.PgxPoolIface, project *models.Project, user *models.User) (pgxmock.PgxPoolIface, error) {
	mockRows := pgxmock.NewRows([]string{"user_id"}).
		AddRow(user.UserID)

	mock.ExpectQuery("SELECT").
		WithArgs(project.ProjectID).
		WillReturnRows(mockRows)

	return mock, nil
}

func DBMockDeleteProjectUsers(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}

func DBMockDeleteUserProjects(mock pgxmock.PgxPoolIface, projectUser *models.ProjectUsers) (pgxmock.PgxPoolIface, error) {
	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.UserID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	return mock, nil
}
