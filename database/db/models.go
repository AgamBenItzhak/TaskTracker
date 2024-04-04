// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Member struct {
	MemberID  int32            `db:"member_id" json:"member_id"`
	Email     string           `db:"email" json:"email"`
	FirstName string           `db:"first_name" json:"first_name"`
	LastName  string           `db:"last_name" json:"last_name"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type MemberCredentials struct {
	MemberID     int32            `db:"member_id" json:"member_id"`
	PasswordHash string           `db:"password_hash" json:"password_hash"`
	PasswordSalt string           `db:"password_salt" json:"password_salt"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Project struct {
	ProjectID   int32            `db:"project_id" json:"project_id"`
	ProjectName string           `db:"project_name" json:"project_name"`
	Description pgtype.Text      `db:"description" json:"description"`
	Status      string           `db:"status" json:"status"`
	StartDate   pgtype.Date      `db:"start_date" json:"start_date"`
	EndDate     pgtype.Date      `db:"end_date" json:"end_date"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type ProjectMember struct {
	MemberID  int32            `db:"member_id" json:"member_id"`
	ProjectID int32            `db:"project_id" json:"project_id"`
	Role      string           `db:"role" json:"role"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Task struct {
	TaskID      int32            `db:"task_id" json:"task_id"`
	TaskGroupID int32            `db:"task_group_id" json:"task_group_id"`
	TaskName    string           `db:"task_name" json:"task_name"`
	Description pgtype.Text      `db:"description" json:"description"`
	Status      string           `db:"status" json:"status"`
	Priority    string           `db:"priority" json:"priority"`
	StartDate   pgtype.Date      `db:"start_date" json:"start_date"`
	EndDate     pgtype.Date      `db:"end_date" json:"end_date"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type TaskGroup struct {
	TaskGroupID int32            `db:"task_group_id" json:"task_group_id"`
	ProjectID   int32            `db:"project_id" json:"project_id"`
	GroupName   string           `db:"group_name" json:"group_name"`
	Description pgtype.Text      `db:"description" json:"description"`
	Status      string           `db:"status" json:"status"`
	Priority    string           `db:"priority" json:"priority"`
	StartDate   pgtype.Date      `db:"start_date" json:"start_date"`
	EndDate     pgtype.Date      `db:"end_date" json:"end_date"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type TaskGroupMember struct {
	MemberID    int32            `db:"member_id" json:"member_id"`
	TaskGroupID int32            `db:"task_group_id" json:"task_group_id"`
	Role        string           `db:"role" json:"role"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type TaskMember struct {
	MemberID  int32            `db:"member_id" json:"member_id"`
	TaskID    int32            `db:"task_id" json:"task_id"`
	Role      string           `db:"role" json:"role"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}