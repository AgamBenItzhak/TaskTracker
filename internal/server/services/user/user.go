package user

import (
	"context"

	"github.com/AgamBenItzhak/TaskTracker/internal/db/models"
	"github.com/AgamBenItzhak/TaskTracker/internal/db/queries"
)

// TODO: Add tests for all functions

func (s *UserService) CreateUser(email, passwordHash, passwordSalt, firstName, lastName string) (models.User, error) {
	ctx := context.Background()
	UserID, err := queries.InsertUser(ctx, s.db.Pool, email, passwordHash, passwordSalt, firstName, lastName)
	if err != nil {
		return models.User{}, err
	}
	return s.GetUserByID(UserID)
}

func (s *UserService) GetUsersIDs() ([]int, error) {
	ctx := context.Background()
	return queries.GetUsersIDs(ctx, s.db.Pool)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	ctx := context.Background()
	users, err := s.GetUsersIDs()
	if err != nil {
		return nil, err
	}
	var usersList []models.User
	for _, userID := range users {
		email, passwordHash, passwordSalt, firstName, lastName, updatedAt, err := queries.GetUserByID(ctx, s.db.Pool, userID)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, models.User{
			UserID:       userID,
			Email:        email,
			PasswordHash: passwordHash,
			PasswordSalt: passwordSalt,
			FirstName:    firstName,
			LastName:     lastName,
			UpdatedAt:    updatedAt,
		})
	}
	return usersList, nil
}

func (s *UserService) GetUserByID(userID int) (models.User, error) {
	ctx := context.Background()
	email, passwordHash, passwordSalt, firstName, lastName, updatedAt, err := queries.GetUserByID(ctx, s.db.Pool, userID)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		Email:        email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		FirstName:    firstName,
		LastName:     lastName,
		UpdatedAt:    updatedAt,
	}, nil
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	ctx := context.Background()
	userID, passwordHash, passwordSalt, firstName, lastName, updatedAt, err := queries.GetUserByEmail(ctx, s.db.Pool, email)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		UserID:       userID,
		Email:        email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		FirstName:    firstName,
		LastName:     lastName,
		UpdatedAt:    updatedAt,
	}, nil
}

func (s *UserService) updateUser(userID int, email, passwordHash, passwordSalt, firstName, lastName string) error {
	ctx := context.Background()
	return queries.UpdateUser(ctx, s.db.Pool, userID, email, passwordHash, passwordSalt, firstName, lastName)
}

func (s *UserService) UpdateUserByID(userID int, email, passwordHash, passwordSalt, firstName, lastName string) error {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	if email == "" {
		email = user.Email
	}

	if passwordHash == "" {
		passwordHash = user.PasswordHash
	}

	if passwordSalt == "" {
		passwordSalt = user.PasswordSalt
	}

	if firstName == "" {
		firstName = user.FirstName
	}

	if lastName == "" {
		lastName = user.LastName
	}

	return s.updateUser(userID, email, passwordHash, passwordSalt, firstName, lastName)
}

func (s *UserService) UpdateUserByEmail(email, passwordHash, passwordSalt, firstName, lastName string) error {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if passwordHash == "" {
		passwordHash = user.PasswordHash
	}

	if passwordSalt == "" {
		passwordSalt = user.PasswordSalt
	}

	if firstName == "" {
		firstName = user.FirstName
	}

	if lastName == "" {
		lastName = user.LastName
	}

	return s.updateUser(user.UserID, email, passwordHash, passwordSalt, firstName, lastName)
}

func (s *UserService) deleteUser(userID int) error {
	ctx := context.Background()
	return queries.DeleteUser(ctx, s.db.Pool, userID)
}

func (s *UserService) DeleteUserByID(userID int) error {
	return s.deleteUser(userID)
}

func (s *UserService) DeleteUserByEmail(email string) error {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return err
	}
	return s.deleteUser(user.UserID)
}
