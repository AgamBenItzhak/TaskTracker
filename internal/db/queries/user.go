package queries

import (
	"context"
)

// InsertUser inserts a new user into the database
func InsertUser(ctx context.Context, dbpool PgxIface, email, passwordHash, passwordSalt, firstName, lastName string) (int, error) {
	var userID int
	err := dbpool.QueryRow(ctx, `
		INSERT INTO users (email, password_hash, password_salt, first_name, last_name)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id
	`, email, passwordHash, passwordSalt, firstName, lastName).Scan(&userID)
	return userID, err
}

// GetUsers retrieves all users from the database
func GetUsersIDs(ctx context.Context, dbpool PgxIface) ([]int, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT user_id
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}

// GetUser retrieves a user from the database
func GetUserByID(ctx context.Context, dbpool PgxIface, userID int) (string, string, string, string, string, string, error) {
	var email, passwordHash, passwordSalt, firstName, lastName, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT email, password_hash, password_salt, first_name, last_name, created_at, updated_at
		FROM users
		WHERE user_id = $1
	`, userID).Scan(&email, &passwordHash, &passwordSalt, &firstName, &lastName, &createdAt, &updatedAt)
	return email, passwordHash, passwordSalt, firstName, lastName, updatedAt, err
}

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(ctx context.Context, dbpool PgxIface, email string) (int, string, string, string, string, string, error) {
	var userID int
	var passwordHash, passwordSalt, firstName, lastName, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT user_id, password_hash, password_salt, first_name, last_name, created_at, updated_at
		FROM users
		WHERE email = $1
	`, email).Scan(&userID, &passwordHash, &passwordSalt, &firstName, &lastName, &createdAt, &updatedAt)
	return userID, passwordHash, passwordSalt, firstName, lastName, updatedAt, err
}

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, dbpool PgxIface, userID int, email, passwordHash, passwordSalt, firstName, lastName string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE users
		SET email = $2, password_hash = $3, password_salt = $4, first_name = $5, last_name = $6, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1
	`, userID, email, passwordHash, passwordSalt, firstName, lastName)
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(ctx context.Context, dbpool PgxIface, userID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM users
		WHERE user_id = $1
	`, userID)
	return err
}

// InsertProjectUser inserts a new project user into the database
func InsertProjectUser(ctx context.Context, dbpool PgxIface, userID, projectID int, role string) error {
	_, err := dbpool.Exec(ctx, `
		INSERT INTO project_users (user_id, project_id, role)
		VALUES ($1, $2, $3)
	`, userID, projectID, role)
	return err
}

// GetProjectUser retrieves a project user from the database
func GetProjectUser(ctx context.Context, dbpool PgxIface, userID, projectID int) (string, string, error) {
	var role, createdAt, updatedAt string
	err := dbpool.QueryRow(ctx, `
		SELECT role, created_at, updated_at
		FROM project_users
		WHERE user_id = $1 AND project_id = $2
	`, userID, projectID).Scan(&role, &createdAt, &updatedAt)
	return role, updatedAt, err
}

// UpdateProjectUser updates a project user in the database
func UpdateProjectUser(ctx context.Context, dbpool PgxIface, userID, projectID int, role string) error {
	_, err := dbpool.Exec(ctx, `
		UPDATE project_users
		SET role = $3, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND project_id = $2
	`, userID, projectID, role)
	return err
}

// DeleteProjectUser deletes a project user from the database
func DeleteProjectUser(ctx context.Context, dbpool PgxIface, userID, projectID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM project_users
		WHERE user_id = $1 AND project_id = $2
	`, userID, projectID)
	return err
}

// GetProjectsByUserID retrieves all projects for a user from the database
func GetProjectsByUserID(ctx context.Context, dbpool PgxIface, userID int) ([]int, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT project_id
		FROM project_users
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projectIDs []int
	for rows.Next() {
		var projectID int
		err := rows.Scan(&projectID)
		if err != nil {
			return nil, err
		}
		projectIDs = append(projectIDs, projectID)
	}
	return projectIDs, nil
}

// GetUsersByProjectID retrieves all users for a project from the database
func GetUsersByProjectID(ctx context.Context, dbpool PgxIface, projectID int) ([]int, error) {
	rows, err := dbpool.Query(ctx, `
		SELECT user_id
		FROM project_users
		WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}

// DeleteProjectUsers deletes all project users from the database for a project
func DeleteProjectUsers(ctx context.Context, dbpool PgxIface, projectID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM project_users
		WHERE project_id = $1
	`, projectID)
	return err
}

// DeleteUserProjects deletes all project users from the database for a user
func DeleteUserProjects(ctx context.Context, dbpool PgxIface, userID int) error {
	_, err := dbpool.Exec(ctx, `
		DELETE FROM project_users
		WHERE user_id = $1
	`, userID)
	return err
}
