// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: member.sql

package db

import (
	"context"
)

const CreateMember = `-- name: CreateMember :one
INSERT INTO member (email, first_name, last_name, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING member_id
`

type CreateMemberParams struct {
	Email     string `db:"email" json:"email"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

// Insert a new member into the database
func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (int32, error) {
	row := q.db.QueryRow(ctx, CreateMember, arg.Email, arg.FirstName, arg.LastName)
	var member_id int32
	err := row.Scan(&member_id)
	return member_id, err
}

const CreateMemberCredentials = `-- name: CreateMemberCredentials :exec
INSERT INTO member_credentials (member_id, password_hash, password_salt, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`

type CreateMemberCredentialsParams struct {
	MemberID     int32  `db:"member_id" json:"member_id"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	PasswordSalt string `db:"password_salt" json:"password_salt"`
}

// Insert a new member's credentials into the database
func (q *Queries) CreateMemberCredentials(ctx context.Context, arg CreateMemberCredentialsParams) error {
	_, err := q.db.Exec(ctx, CreateMemberCredentials, arg.MemberID, arg.PasswordHash, arg.PasswordSalt)
	return err
}

const DeleteMemberByID = `-- name: DeleteMemberByID :exec
DELETE FROM member WHERE member_id = $1
`

// Delete a member from the database
func (q *Queries) DeleteMemberByID(ctx context.Context, memberID int32) error {
	_, err := q.db.Exec(ctx, DeleteMemberByID, memberID)
	return err
}

const DeleteMemberCredentialsByID = `-- name: DeleteMemberCredentialsByID :exec
DELETE FROM member_credentials WHERE member_id = $1
`

// Delete a member's credentials from the database
func (q *Queries) DeleteMemberCredentialsByID(ctx context.Context, memberID int32) error {
	_, err := q.db.Exec(ctx, DeleteMemberCredentialsByID, memberID)
	return err
}

const GetAllMemberIDs = `-- name: GetAllMemberIDs :many
SELECT member_id FROM member
`

// Get all member IDs from the database
func (q *Queries) GetAllMemberIDs(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, GetAllMemberIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var member_id int32
		if err := rows.Scan(&member_id); err != nil {
			return nil, err
		}
		items = append(items, member_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAllMemberProjectIDs = `-- name: GetAllMemberProjectIDs :many
SELECT project_id FROM project_member WHERE member_id = $1
`

// Get all project IDs assigned to a member
func (q *Queries) GetAllMemberProjectIDs(ctx context.Context, memberID int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, GetAllMemberProjectIDs, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var project_id int32
		if err := rows.Scan(&project_id); err != nil {
			return nil, err
		}
		items = append(items, project_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAllMemberProjects = `-- name: GetAllMemberProjects :many
SELECT member_id, project_id, role, created_at, updated_at FROM project_member WHERE member_id = $1
`

// Get all projects assigned to a member
func (q *Queries) GetAllMemberProjects(ctx context.Context, memberID int32) ([]ProjectMember, error) {
	rows, err := q.db.Query(ctx, GetAllMemberProjects, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectMember
	for rows.Next() {
		var i ProjectMember
		if err := rows.Scan(
			&i.MemberID,
			&i.ProjectID,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAllMembers = `-- name: GetAllMembers :many
SELECT member_id, email, first_name, last_name, created_at, updated_at FROM member
`

// Get all members from the database
func (q *Queries) GetAllMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.Query(ctx, GetAllMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Member
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.MemberID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetMemberByID = `-- name: GetMemberByID :one
SELECT member_id, email, first_name, last_name, created_at, updated_at FROM member WHERE member_id = $1
`

// Get a member by their ID
func (q *Queries) GetMemberByID(ctx context.Context, memberID int32) (Member, error) {
	row := q.db.QueryRow(ctx, GetMemberByID, memberID)
	var i Member
	err := row.Scan(
		&i.MemberID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetMemberCredentialsByID = `-- name: GetMemberCredentialsByID :one
SELECT member_id, password_hash, password_salt, created_at, updated_at FROM member_credentials WHERE member_id = $1
`

// Get a member's credentials by their ID
func (q *Queries) GetMemberCredentialsByID(ctx context.Context, memberID int32) (MemberCredentials, error) {
	row := q.db.QueryRow(ctx, GetMemberCredentialsByID, memberID)
	var i MemberCredentials
	err := row.Scan(
		&i.MemberID,
		&i.PasswordHash,
		&i.PasswordSalt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const UpdateMemberByID = `-- name: UpdateMemberByID :exec
UPDATE member SET email = $2, first_name = $3, last_name = $4, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1
`

type UpdateMemberByIDParams struct {
	MemberID  int32  `db:"member_id" json:"member_id"`
	Email     string `db:"email" json:"email"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

// Update a member's information
func (q *Queries) UpdateMemberByID(ctx context.Context, arg UpdateMemberByIDParams) error {
	_, err := q.db.Exec(ctx, UpdateMemberByID,
		arg.MemberID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	return err
}

const UpdateMemberCredentialsByID = `-- name: UpdateMemberCredentialsByID :exec
UPDATE member_credentials SET password_hash = $2, password_salt = $3, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1
`

type UpdateMemberCredentialsByIDParams struct {
	MemberID     int32  `db:"member_id" json:"member_id"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	PasswordSalt string `db:"password_salt" json:"password_salt"`
}

// Update a member's credentials
func (q *Queries) UpdateMemberCredentialsByID(ctx context.Context, arg UpdateMemberCredentialsByIDParams) error {
	_, err := q.db.Exec(ctx, UpdateMemberCredentialsByID, arg.MemberID, arg.PasswordHash, arg.PasswordSalt)
	return err
}