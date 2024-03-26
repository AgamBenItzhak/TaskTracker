package user

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
)

// UserService provides user-related functions
type UserService struct {
	db *db.DB
}

// NewUserService creates a new UserService instance
func NewUserService(db *db.DB) *UserService {
	return &UserService{db: db}
}
