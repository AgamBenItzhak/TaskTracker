package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries/testmockdb"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockInsertTasksGroup(mock, newTasksGroup)
	require.NoError(t, err)

	tasksGroupID, err := InsertTasksGroup(context.Background(), mock, newTasksGroup)
	require.NoError(t, err)
	require.Equal(t, newTasksGroup.TaskGroupID, tasksGroupID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAllTasksGroupsIDsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockGetAllTasksGroupsIDsByProjectID(mock, newTasksGroup)
	require.NoError(t, err)

	tasksGroupIDs, err := GetAllTasksGroupsIDsByProjectID(context.Background(), mock, newTasksGroup.ProjectID)
	require.NoError(t, err)

	require.Equal(t, newTasksGroup.TaskGroupID, tasksGroupIDs[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAllTasksGroupsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockGetAllTasksGroupsByProjectID(mock, newTasksGroup)
	require.NoError(t, err)

	tasksGroups, err := GetAllTasksGroupsByProjectID(context.Background(), mock, newTasksGroup.ProjectID)
	require.NoError(t, err)

	require.Equal(t, newTasksGroup, tasksGroups[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetTasksGroupsByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockGetTasksGroupsByID(mock, newTasksGroup)
	require.NoError(t, err)

	tasksGroup, err := GetTasksGroupsByID(context.Background(), mock, newTasksGroup.TaskGroupID)
	require.NoError(t, err)

	require.Equal(t, newTasksGroup, tasksGroup)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockUpdateTasksGroup(mock, newTasksGroup)
	require.NoError(t, err)

	err = UpdateTasksGroup(context.Background(), mock, newTasksGroup)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockDeleteTasksGroup(mock, TasksGroup)
	require.NoError(t, err)

	err = DeleteTasksGroup(context.Background(), mock, TasksGroup.TaskGroupID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteAllTasksGroupsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	project := models.NewMockProject()

	mock, err = testmockdb.DBMockDeleteAllTasksGroupsByProjectID(mock, project)
	require.NoError(t, err)

	err = DeleteAllTasksGroupsByProjectID(context.Background(), mock, project.ProjectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTask := models.NewMockTasks()

	mock, err = testmockdb.DBMockInsertTask(mock, newTask)
	require.NoError(t, err)

	taskID, err := InsertTask(context.Background(), mock, newTask)
	require.NoError(t, err)
	require.Equal(t, 1, taskID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetTaskByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTask := models.NewMockTasks()

	mock, err = testmockdb.DBMockGetTaskByID(mock, newTask)
	require.NoError(t, err)

	task, err := GetTaskByID(context.Background(), mock, newTask.TaskID)
	require.NoError(t, err)

	require.Equal(t, newTask, task)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAllTasksByTasksGroupID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTask := models.NewMockTasks()

	mock, err = testmockdb.DBMockGetAllTasksByTasksGroupID(mock, newTask)
	require.NoError(t, err)

	tasks, err := GetAllTasksByTasksGroupID(context.Background(), mock, newTask.TaskGroupID)
	require.NoError(t, err)

	require.Equal(t, newTask, tasks[0])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateTaskByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTask := models.NewMockTasks()

	mock, err = testmockdb.DBMockUpdateTaskByID(mock, newTask)
	require.NoError(t, err)

	err = UpdateTaskByID(context.Background(), mock, newTask)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteAllTasksByTasksGroupID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroup := models.NewMockTasksGroups()

	mock, err = testmockdb.DBMockDeleteAllTasksByTasksGroupID(mock, TasksGroup)
	require.NoError(t, err)

	err = DeleteAllTasksByTasksGroupID(context.Background(), mock, TasksGroup.TaskGroupID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTaskByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	task := models.NewMockTasks()

	mock, err = testmockdb.DBMockDeleteTaskByID(mock, task)
	require.NoError(t, err)

	err = DeleteTaskByID(context.Background(), mock, task.TaskID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
