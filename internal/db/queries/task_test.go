package queries

import (
	"context"
	"testing"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func TestInsertTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mock.ExpectQuery("INSERT INTO tasks_groups").
		WithArgs(newTasksGroup.ProjectID, newTasksGroup.GroupName, newTasksGroup.Description).
		WillReturnRows(pgxmock.NewRows([]string{"tasks_group_id"}).AddRow(1))

	tasksGroupID, err := InsertTasksGroup(context.Background(), mock, newTasksGroup)
	require.NoError(t, err)
	require.Equal(t, 1, tasksGroupID)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAllTasksGroupsIDsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTasksGroup := models.NewMockTasksGroups()

	mockRows := pgxmock.NewRows([]string{"tasks_group_id"}).
		AddRow(newTasksGroup.TaskGroupID)

	mock.ExpectQuery("SELECT").
		WithArgs(newTasksGroup.ProjectID).
		WillReturnRows(mockRows)

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

	mockRows := pgxmock.NewRows([]string{"task_group_id", "project_id", "group_name", "description", "created_at", "updated_at"}).
		AddRow(newTasksGroup.TaskGroupID, newTasksGroup.ProjectID, newTasksGroup.GroupName, newTasksGroup.Description, newTasksGroup.CreatedAt, newTasksGroup.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newTasksGroup.ProjectID).
		WillReturnRows(mockRows)

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

	mockRows := pgxmock.NewRows([]string{"task_group_id", "project_id", "group_name", "description", "created_at", "updated_at"}).
		AddRow(newTasksGroup.TaskGroupID, newTasksGroup.ProjectID, newTasksGroup.GroupName, newTasksGroup.Description, newTasksGroup.CreatedAt, newTasksGroup.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newTasksGroup.TaskGroupID).
		WillReturnRows(mockRows)

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

	mock.ExpectExec("UPDATE tasks_groups").
		WithArgs(newTasksGroup.TaskGroupID, newTasksGroup.ProjectID, newTasksGroup.GroupName, newTasksGroup.Description).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateTasksGroup(context.Background(), mock, newTasksGroup)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTasksGroup(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	TasksGroupID := 1

	mock.ExpectExec("DELETE FROM tasks_groups").
		WithArgs(TasksGroupID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteTasksGroup(context.Background(), mock, TasksGroupID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteAllTasksGroupsByProjectID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	projectID := 1

	mock.ExpectExec("DELETE FROM tasks_groups").
		WithArgs(projectID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteAllTasksGroupsByProjectID(context.Background(), mock, projectID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsertTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	newTask := models.NewMockTasks()

	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(newTask.TaskGroupID, newTask.TaskName, newTask.Description, newTask.Status, newTask.Priority, newTask.StartDate, newTask.EndDate).
		WillReturnRows(pgxmock.NewRows([]string{"task_id"}).AddRow(1))

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

	mockRows := pgxmock.NewRows([]string{"task_id", "tasks_group_id", "task_name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(newTask.TaskID, newTask.TaskGroupID, newTask.TaskName, newTask.Description, newTask.Status, newTask.Priority, newTask.StartDate, newTask.EndDate, newTask.CreatedAt, newTask.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newTask.TaskID).
		WillReturnRows(mockRows)

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

	mockRows := pgxmock.NewRows([]string{"task_id", "task_group_id", "task_name", "description", "status", "priority", "start_date", "end_date", "created_at", "updated_at"}).
		AddRow(newTask.TaskID, newTask.TaskGroupID, newTask.TaskName, newTask.Description, newTask.Status, newTask.Priority, newTask.StartDate, newTask.EndDate, newTask.CreatedAt, newTask.UpdatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(newTask.TaskGroupID).
		WillReturnRows(mockRows)

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

	mock.ExpectExec("UPDATE tasks").
		WithArgs(newTask.TaskID, newTask.TaskGroupID, newTask.TaskName, newTask.Description, newTask.Status, newTask.Priority, newTask.StartDate, newTask.EndDate).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = UpdateTaskByID(context.Background(), mock, newTask)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteTaskByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	taskID := 1

	mock.ExpectExec("DELETE FROM tasks").
		WithArgs(taskID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = DeleteTaskByID(context.Background(), mock, taskID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
