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