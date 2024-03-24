package queries

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	email := "test@example.com"
	passwordHash := "passwordHash"
	passwordSalt := "passwordSalt"
	firstName := "firstName"
	lastName := "lastName"

	mock.ExpectExec("INSERT INTO users").
		WithArgs(email, passwordHash, passwordSalt, firstName, lastName).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = InsertUser(context.Background(), mock, email, passwordHash, passwordSalt, firstName, lastName)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mockRows := pgxmock.NewRows([]string{"user_id"}).
		AddRow(1).
		AddRow(2)

	mock.ExpectQuery("SELECT").
		WillReturnRows(mockRows)

	userIDs, err := GetUsers(context.Background(), mock)
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

	userID := 1
	mockRows := pgxmock.NewRows([]string{"email", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at"}).
		AddRow("test@example.com", "passwordHash", "passwordSalt", "firstName", "lastName", "2021-01-01T00:00:00Z", "2021-01-01T00:00:00Z")

	mock.ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnRows(mockRows)

	email, passwordHash, passwordSalt, firstName, lastName, updatedAt, err := GetUserByID(context.Background(), mock, userID)
	require.NoError(t, err)

	require.Equal(t, "test@example.com", email)
	require.Equal(t, "passwordHash", passwordHash)
	require.Equal(t, "passwordSalt", passwordSalt)
	require.Equal(t, "firstName", firstName)
	require.Equal(t, "lastName", lastName)
	require.Equal(t, "2021-01-01T00:00:00Z", updatedAt)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	// Create a mock connection
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	email := "test@example.com"
	mockRows := pgxmock.NewRows([]string{"user_id", "password_hash", "password_salt", "first_name", "last_name", "created_at", "updated_at"}).
		AddRow(1, "passwordHash", "passwordSalt", "firstName", "lastName", "2021-01-01T00:00:00Z", "2021-01-01T00:00:00Z")

	// Expect a SELECT query with the given email argument and return the mock rows
	mock.ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(mockRows)

	// Call the GetUserByEmail function and retrieve the returned values
	userID, passwordHash, passwordSalt, firstName, lastName, updatedAt, err := GetUserByEmail(context.Background(), mock, email)
	require.NoError(t, err)

	// Assert the expected values
	require.Equal(t, 1, userID)
	require.Equal(t, "passwordHash", passwordHash)
	require.Equal(t, "passwordSalt", passwordSalt)
	require.Equal(t, "firstName", firstName)
	require.Equal(t, "lastName", lastName)
	require.Equal(t, "2021-01-01T00:00:00Z", updatedAt)

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	email := "test@example.com"
	passwordHash := "passwordHash"
	passwordSalt := "passwordSalt"
	firstName := "firstName"
	lastName := "lastName"

	mock.ExpectExec("UPDATE users").
		WithArgs(userID, email, passwordHash, passwordSalt, firstName, lastName).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateUser(context.Background(), mock, userID, email, passwordHash, passwordSalt, firstName, lastName)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1

	mock.ExpectExec("DELETE FROM users").
		WithArgs(userID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteUser(context.Background(), mock, userID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1
	role := "admin"

	mock.ExpectExec("INSERT INTO project_users").
		WithArgs(userID, projectID, role).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = InsertProjectUser(context.Background(), mock, userID, projectID, role)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1
	role := "admin"

	mockRows := pgxmock.NewRows([]string{"role", "created_at", "updated_at"}).
		AddRow(role, "2021-01-01T00:00:00Z", "2021-01-01T00:00:00Z")

	mock.ExpectQuery("SELECT").
		WithArgs(userID, projectID).
		WillReturnRows(mockRows)

	retrievedRole, updatedAt, err := GetProjectUser(context.Background(), mock, userID, projectID)
	require.NoError(t, err)

	require.Equal(t, role, retrievedRole)
	require.Equal(t, "2021-01-01T00:00:00Z", updatedAt)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1
	role := "admin"

	mock.ExpectExec("UPDATE project_users").
		WithArgs(userID, projectID, role).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateProjectUser(context.Background(), mock, userID, projectID, role)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(userID, projectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProjectUser(context.Background(), mock, userID, projectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetProjectsByUserID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1

	mockRows := pgxmock.NewRows([]string{"project_id"}).
		AddRow(projectID)

	mock.ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnRows(mockRows)

	projectIDs, err := GetProjectsByUserID(context.Background(), mock, userID)
	require.NoError(t, err)

	require.Len(t, projectIDs, 1)
	require.Equal(t, projectID, projectIDs[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsersByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1
	projectID := 1

	mockRows := pgxmock.NewRows([]string{"user_id"}).
		AddRow(userID)

	mock.ExpectQuery("SELECT").
		WithArgs(projectID).
		WillReturnRows(mockRows)

	userIDs, err := GetUsersByProjectID(context.Background(), mock, projectID)
	require.NoError(t, err)

	require.Len(t, userIDs, 1)
	require.Equal(t, userID, userIDs[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteProjectUsers(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectID := 1

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(projectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteProjectUsers(context.Background(), mock, projectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUserProjects(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	userID := 1

	mock.ExpectExec("DELETE FROM project_users").
		WithArgs(userID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteUserProjects(context.Background(), mock, userID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
