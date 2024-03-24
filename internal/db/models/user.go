package models

type User struct {
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	LastSeen     string `json:"last_seen"`
}

type ProjectUsers struct {
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
