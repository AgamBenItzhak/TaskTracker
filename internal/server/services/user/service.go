package user

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/pashagolub/pgxmock/v3"
)

// UserService provides user-related functions
type UserService struct {
	dbpool db.PgxIface
}

// NewUserService creates a new UserService instance
func NewUserService(dbpool db.PgxIface) *UserService {
	return &UserService{dbpool: dbpool}
}

func NewMockUserService() *UserService {
	mock, _ := pgxmock.NewPool()
	return NewUserService(mock)
}
