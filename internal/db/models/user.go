package models

type User struct {
	UserID       int    `validate:"required" json:"user_id" mapstructure:"user_id"`
	Email        string `validate:"required,email" json:"email" mapstructure:"email"`
	PasswordHash string `validate:"required" json:"password_hash" mapstructure:"password_hash"`
	PasswordSalt string `validate:"required" json:"password_salt" mapstructure:"password_salt"`
	FirstName    string `validate:"required" json:"first_name" mapstructure:"first_name"`
	LastName     string `validate:"required" json:"last_name" mapstructure:"last_name"`
	CreatedAt    string `validate:"required,date" json:"created_at" mapstructure:"created_at"`
	UpdatedAt    string `validate:"required,date" json:"updated_at" mapstructure:"updated_at"`
	LastSeen     string `validate:"required,date" json:"last_seen" mapstructure:"last_seen"`
}

type ProjectUsers struct {
	UserID    int    `validate:"required" json:"user_id" mapstructure:"user_id"`
	ProjectID int    `validate:"required" json:"project_id" mapstructure:"project_id"`
	Role      string `validate:"required" json:"role" mapstructure:"role"`
	CreatedAt string `validate:"required,date" json:"created_at" mapstructure:"created_at"`
	UpdatedAt string `validate:"required,date" json:"updated_at" mapstructure:"updated_at"`
}

func NewUser(userID int, email, passwordHash, passwordSalt, firstName, lastName, createdAt, updatedAt, lastSeen string) *User {
	return &User{
		UserID:       userID,
		Email:        email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		FirstName:    firstName,
		LastName:     lastName,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		LastSeen:     lastSeen,
	}
}

func NewMockUser() *User {
	return &User{
		UserID:       1,
		Email:        "test@example.com",
		PasswordHash: "passwordHash",
		PasswordSalt: "passwordSalt",
		FirstName:    "firstName",
		LastName:     "lastName",
		CreatedAt:    "2021-01-01T00:00:00Z",
		UpdatedAt:    "2021-01-01T00:00:00Z",
		LastSeen:     "2021-01-01T00:00:00Z",
	}
}

func NewProjectUsers(userID, projectID int, role, createdAt, updatedAt string) *ProjectUsers {
	return &ProjectUsers{
		UserID:    userID,
		ProjectID: projectID,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewMockProjectUsers() *ProjectUsers {
	return &ProjectUsers{
		UserID:    1,
		ProjectID: 1,
		Role:      "role",
		CreatedAt: "2021-01-01T00:00:00Z",
		UpdatedAt: "2021-01-01T00:00:00Z",
	}
}
