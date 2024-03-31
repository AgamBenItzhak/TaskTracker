package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	UserID    int       `validate:"required" json:"user_id" mapstructure:"user_id"`
	FirstName string    `validate:"required" json:"first_name" mapstructure:"first_name"`
	LastName  string    `validate:"required" json:"last_name" mapstructure:"last_name"`
	CreatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
	LastSeen  time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"last_seen" mapstructure:"last_seen"`
}

type UserCredentials struct {
	UserID       int       `validate:"required" json:"user_id" mapstructure:"user_id"`
	Email        string    `validate:"required,email" json:"email" mapstructure:"email"`
	PasswordHash string    `validate:"required" json:"password_hash" mapstructure:"password_hash"`
	PasswordSalt string    `validate:"required" json:"password_salt" mapstructure:"password_salt"`
	CreatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"created_at" mapstructure:"created_at"`
	UpdatedAt    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"updated_at" mapstructure:"updated_at"`
}

func NewUser(userID int, firstName, lastName string, createdAt, updatedAt, lastSeen time.Time) *User {
	return &User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		LastSeen:  lastSeen,
	}
}

func NewMockUser() *User {
	return &User{
		UserID:    1,
		FirstName: "firstName",
		LastName:  "lastName",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LastSeen:  time.Now(),
	}
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func NewUserCredentials(userID int, email, passwordHash, passwordSalt string, createdAt, updatedAt time.Time) *UserCredentials {
	return &UserCredentials{
		UserID:       userID,
		Email:        email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func NewMockUserCredentials() *UserCredentials {
	return &UserCredentials{
		UserID:       1,
		Email:        "email",
		PasswordHash: "passwordHash",
		PasswordSalt: "passwordSalt",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (u *UserCredentials) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
