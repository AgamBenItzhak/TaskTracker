package task

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries"
)

func (s *TaskService) CreateTasksGroups(tasksGroup *models.TasksGroups) (int, error) {
	ctx := context.Background()
	return queries.InsertTasksGroup(ctx, s.dbpool, tasksGroup)
}

func (s *TaskService) GetAllTasksGroupsIDsByProjectID(projectID int) ([]int, error) {
	ctx := context.Background()
	return queries.GetAllTasksGroupsIDsByProjectID(ctx, s.dbpool, projectID)
}

func (s *TaskService) GetAllTasksGroupsByProjectID(projectID int) ([]*models.TasksGroups, error) {
	ctx := context.Background()
	return queries.GetAllTasksGroupsByProjectID(ctx, s.dbpool, projectID)
}

func (s *TaskService) GetTasksGroupsByID(TasksGroupID int) (*models.TasksGroups, error) {
	ctx := context.Background()
	return queries.GetTasksGroupsByID(ctx, s.dbpool, TasksGroupID)
}

func (s *TaskService) UpdateTasksGroup(TasksGroup *models.TasksGroups) error {
	ctx := context.Background()
	return queries.UpdateTasksGroup(ctx, s.dbpool, TasksGroup)
}

func (s *TaskService) DeleteTasksGroup(TasksGroupID int) error {
	ctx := context.Background()
	return queries.DeleteTasksGroup(ctx, s.dbpool, TasksGroupID)
}

func (s *TaskService) DeleteAllTasksGroupsByProjectID(projectID int) error {
	ctx := context.Background()
	return queries.DeleteAllTasksGroupsByProjectID(ctx, s.dbpool, projectID)
}

func (s *TaskService) CreateTask(taskGroupID int, taskName, description, status, priority, startDate, endDate string) (int, error) {
	ctx := context.Background()
	task := models.NewTasks(0, taskGroupID, taskName, description, status, priority, startDate, endDate, "", "")
	return queries.InsertTask(ctx, s.dbpool, task)
}

func (s *TaskService) GetAllTasksByTasksGroupID(TasksGroupID int) ([]*models.Tasks, error) {
	ctx := context.Background()
	return queries.GetAllTasksByTasksGroupID(ctx, s.dbpool, TasksGroupID)
}

func (s *TaskService) GetTaskByID(taskID int) (*models.Tasks, error) {
	ctx := context.Background()
	return queries.GetTaskByID(ctx, s.dbpool, taskID)
}

func (s *TaskService) UpdateTaskByID(task *models.Tasks) error {
	ctx := context.Background()
	return queries.UpdateTaskByID(ctx, s.dbpool, task)
}

func (s *TaskService) DeleteAllTasksByTasksGroupID(TasksGroupID int) error {
	ctx := context.Background()
	return queries.DeleteAllTasksByTasksGroupID(ctx, s.dbpool, TasksGroupID)
}

func (s *TaskService) DeleteTaskByID(taskID int) error {
	ctx := context.Background()
	return queries.DeleteTaskByID(ctx, s.dbpool, taskID)
}
