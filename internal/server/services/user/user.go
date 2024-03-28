package user

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries"
)

// TODO: Add tests for all functions

func (s *UserService) CreateUser(email, password, firstName, lastName string) (*models.User, error) {
	ctx := context.Background()
	userData := models.NewUser(0, email, password, firstName, lastName, "", "", "", "")
	userID, err := queries.InsertUser(ctx, s.dbpool, userData)
	if err != nil {
		return nil, err
	}

	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUsersIDs() ([]int, error) {
	ctx := context.Background()
	return queries.GetUsersIDs(ctx, s.dbpool)
}

func (s *UserService) GetUsers() ([]*models.User, error) {
	ctx := context.Background()
	users, err := queries.GetUsers(ctx, s.dbpool)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	ctx := context.Background()
	user, err := queries.GetUserByID(ctx, s.dbpool, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	user, err := queries.GetUserByEmail(ctx, s.dbpool, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	ctx := context.Background()
	return queries.UpdateUser(ctx, s.dbpool, user)
}

func (s *UserService) DeleteUserByID(userID int) error {
	return queries.DeleteUserByID(context.Background(), s.dbpool, userID)
}

func (s *UserService) DeleteUserByEmail(email string) error {
	return queries.DeleteUserByEmail(context.Background(), s.dbpool, email)
}
