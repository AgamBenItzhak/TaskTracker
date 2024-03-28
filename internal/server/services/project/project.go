package project

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries"
)

func (s *ProjectService) CreateProject(projectName, description, status, startDate, endDate string) (*models.Project, error) {
	ctx := context.Background()
	project := models.NewProject(0, projectName, description, status, startDate, endDate, "", "")
	projectID, err := queries.InsertProject(ctx, s.dbpool, project)
	if err != nil {
		return nil, err
	}

	newProject, err := s.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}

func (ps *ProjectService) GetProjectsIDs() ([]int, error) {
	ctx := context.Background()
	return queries.GetProjectsIDs(ctx, ps.dbpool)
}

func (ps *ProjectService) GetProjects() ([]*models.Project, error) {
	projects, err := ps.GetProjectsIDs()
	if err != nil {
		return nil, err
	}

	var projectsList []*models.Project
	for _, projectID := range projects {
		project, err := ps.GetProjectByID(projectID)
		if err != nil {
			return nil, err
		}
		projectsList = append(projectsList, project)
	}

	return projectsList, nil
}

func (ps *ProjectService) GetProjectByID(projectID int) (*models.Project, error) {
	ctx := context.Background()
	return queries.GetProjectByID(ctx, ps.dbpool, projectID)
}

func (ps *ProjectService) UpdateProject(project *models.Project) (*models.Project, error) {
	ctx := context.Background()
	err := queries.UpdateProject(ctx, ps.dbpool, project)
	if err != nil {
		return nil, err
	}

	updatedProject, err := ps.GetProjectByID(project.ProjectID)
	if err != nil {
		return nil, err
	}

	return updatedProject, nil
}
