package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mockRows := pgxmock.NewRows([]string{"user_id"}).AddRow(newUser.UserID)

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(newUser.Email, newUser.PasswordHash, newUser.PasswordSalt, newUser.FirstName, newUser.LastName).
		WillReturnRows(mockRows)

	userID, err := InsertUser(context.Background(), mock, newUser)
	require.NoError(t, err)

	require.Equal(t, newUser.UserID, userID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsersIDs(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mockRows := pgxmock.NewRows([]string{"user_id"}).
		AddRow(1).
		AddRow(2)

	mock.ExpectQuery("SELECT").
		WillReturnRows(mockRows)

	userIDs, err := GetUsersIDs(context.Background(), mock)
	require.NoError(t, err)

	require.Len(t, userIDs, 2)
	require.Equal(t, 1, userIDs[0])
	require.Equal(t, 2, userIDs[1])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mockRows := pgxmock.NewRows([]string{"user_id", "email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at", "last_seen"}).
		AddRow(newUser.UserID, newUser.Email, newUser.PasswordHash, newUser.PasswordSalt, newUser.FirstName, newUser.LastName, newUser.CreatedAt, newUser.UpdatedAt, newUser.LastSeen)

	mock.ExpectQuery("SELECT").
		WithArgs(newUser.UserID).
		WillReturnRows(mockRows)

	user, err := GetUserByID(context.Background(), mock, newUser.UserID)
	require.NoError(t, err)

	require.Equal(t, newUser, user)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	// Create a mock connection
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	// Create a mock row
	mockRows := pgxmock.NewRows([]string{"user_id", "email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at", "last_seen"}).
		AddRow(newUser.UserID, newUser.Email, newUser.PasswordHash, newUser.PasswordSalt, newUser.FirstName, newUser.LastName, newUser.CreatedAt, newUser.UpdatedAt, newUser.LastSeen)

	// Expect a query to be executed
	mock.ExpectQuery("SELECT").
		WithArgs(newUser.Email).
		WillReturnRows(mockRows)

	// Call the function we are testing
	user, err := GetUserByEmail(context.Background(), mock, newUser.Email)
	require.NoError(t, err)

	// Check if the returned user is the same as the one we created
	require.Equal(t, newUser, user)

	// Check if all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mock.ExpectExec("UPDATE users").
		WithArgs(newUser.UserID, newUser.Email, newUser.PasswordHash, newUser.PasswordSalt, newUser.FirstName, newUser.LastName).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateUser(context.Background(), mock, newUser)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	user := models.NewMockUser()

	mock.ExpectExec("DELETE FROM users").
		WithArgs(user.UserID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteUser(context.Background(), mock, user.UserID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock.ExpectExec("INSERT INTO project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID, projectUser.Role).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = InsertProjectUser(context.Background(), mock, projectUser.UserID, projectUser.ProjectID, projectUser.Role)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newProjectUser := models.NewMockProjectUsers()

	mockRows := pgxmock.NewRows([]string{"user_id", "project_id", "role", "created_at", "updated_at"}).
		AddRow(newProjectUser.UserID, newProjectUser.ProjectID, newProjectUser.Role, newProjectUser.CreatedAt, newProjectUser.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newProjectUser.UserID, newProjectUser.ProjectID).
		WillReturnRows(mockRows)

	projectUser, err := GetProjectUser(context.Background(), mock, newProjectUser.UserID, newProjectUser.ProjectID)
	require.NoError(t, err)

	require.Equal(t, newProjectUser, projectUser)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock.ExpectExec("UPDATE project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID, projectUser.Role).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateProjectUser(context.Background(), mock, projectUser.UserID, projectUser.ProjectID, projectUser.Role)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.UserID, projectUser.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProjectUser(context.Background(), mock, projectUser.UserID, projectUser.ProjectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectsByUserID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	user := models.NewMockUser()
	project := models.NewMockProject()

	mockRows := pgxmock.NewRows([]string{"project_id"}).
		AddRow(project.ProjectID)

	mock.ExpectQuery("SELECT").
		WithArgs(user.UserID).
		WillReturnRows(mockRows)

	projectIDs, err := GetProjectsByUserID(context.Background(), mock, user.UserID)
	require.NoError(t, err)

	require.Len(t, projectIDs, 1)
	require.Equal(t, project.ProjectID, projectIDs[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsersByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	user := models.NewMockUser()
	project := models.NewMockProject()

	mockRows := pgxmock.NewRows([]string{"user_id"}).
		AddRow(user.UserID)

	mock.ExpectQuery("SELECT").
		WithArgs(project.ProjectID).
		WillReturnRows(mockRows)

	userIDs, err := GetUsersByProjectID(context.Background(), mock, project.ProjectID)
	require.NoError(t, err)

	require.Len(t, userIDs, 1)
	require.Equal(t, user.UserID, userIDs[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProjectUsers(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.ProjectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProjectUsers(context.Background(), mock, projectUser.ProjectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUserProjects(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectUser.UserID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteUserProjects(context.Background(), mock, projectUser.UserID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
