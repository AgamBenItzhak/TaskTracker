-- Insert a new member into the database
-- name: CreateMember :one
INSERT INTO member (email, first_name, last_name, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING member_id;

-- Get all member IDs from the database
-- name: GetAllMemberIDs :many
SELECT member_id FROM member;

-- Get all members from the database
-- name: GetAllMembers :many
SELECT * FROM member;

-- Get a member by their ID
-- name: GetMemberByID :one
SELECT * FROM member WHERE member_id = $1;

-- Update a member's information
-- name: UpdateMemberByID :exec
UPDATE member SET email = $2, first_name = $3, last_name = $4, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1;

-- Delete a member from the database
-- name: DeleteMemberByID :exec
DELETE FROM member WHERE member_id = $1;

-- Insert a new member's credentials into the database
-- name: CreateMemberCredentials :exec
INSERT INTO member_credentials (member_id, password_hash, password_salt, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Get a member's credentials by their ID
-- name: GetMemberCredentialsByID :one
SELECT * FROM member_credentials WHERE member_id = $1;

-- Update a member's credentials
-- name: UpdateMemberCredentialsByID :exec
UPDATE member_credentials SET password_hash = $2, password_salt = $3, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1;

-- Delete a member's credentials from the database
-- name: DeleteMemberCredentialsByID :exec
DELETE FROM member_credentials WHERE member_id = $1;