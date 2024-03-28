package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries/testmockdb"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mock, err = testmockdb.DBMockInsertUser(mock, newUser)
	require.NoError(t, err)

	userID, err := InsertUser(context.Background(), mock, newUser)
	require.NoError(t, err)

	require.Equal(t, newUser.UserID, userID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mock, err = testmockdb.DBMockGetUsers(mock, newUser)
	require.NoError(t, err)

	users, err := GetUsers(context.Background(), mock)
	require.NoError(t, err)

	require.Len(t, users, 1)
	require.Equal(t, newUser, users[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUsersIDs(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	var users []*models.User
	users = append(users, models.NewMockUser())
	users = append(users, models.NewMockUser())

	mock, err = testmockdb.DBMockGetUsersIDs(mock, users)
	require.NoError(t, err)

	userIDs, err := GetUsersIDs(context.Background(), mock)
	require.NoError(t, err)

	require.Len(t, userIDs, 2)
	require.Equal(t, 1, userIDs[0])
	require.Equal(t, 1, userIDs[1])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newUser := models.NewMockUser()

	mock, err = testmockdb.DBMockGetUserByID(mock, newUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockGetUserByEmail(mock, newUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockUpdateUser(mock, newUser)
	require.NoError(t, err)

	err = UpdateUser(context.Background(), mock, newUser)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUserByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	user := models.NewMockUser()

	mock, err = testmockdb.DBMockDeleteUserByID(mock, user)
	require.NoError(t, err)

	err = DeleteUserByID(context.Background(), mock, user.UserID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteUserByEmail(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	user := models.NewMockUser()

	mock, err = testmockdb.DBMockDeleteUserByEmail(mock, user)
	require.NoError(t, err)

	err = DeleteUserByEmail(context.Background(), mock, user.Email)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertProjectUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectUser := models.NewMockProjectUsers()

	mock, err = testmockdb.DBMockInsertProjectUser(mock, projectUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockGetProjectUser(mock, newProjectUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockUpdateProjectUser(mock, projectUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockDeleteProjectUser(mock, projectUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockGetProjectsByUserID(mock, user, project)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockGetUsersByProjectID(mock, project, user)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockDeleteProjectUsers(mock, projectUser)
	require.NoError(t, err)

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

	mock, err = testmockdb.DBMockDeleteUserProjects(mock, projectUser)
	require.NoError(t, err)

	err = DeleteUserProjects(context.Background(), mock, projectUser.UserID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
