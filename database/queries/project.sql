-- Insert a new project into the database
-- name: CreateProject :one
INSERT INTO project (project_name, description, status, start_date, end_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING project_id;

-- Get all projects from the database
-- name: GetAllProjects :many
SELECT * FROM project;

-- Get a project by its ID
-- name: GetProjectByID :one
SELECT * FROM project WHERE project_id = $1;

-- Update a project's information
-- name: UpdateProjectByID :exec
UPDATE project SET project_name = $2, description = $3, status = $4, start_date = $5, end_date = $6, updated_at = CURRENT_TIMESTAMP
WHERE project_id = $1;

-- Delete a project from the database
-- name: DeleteProjectByID :exec
DELETE FROM project WHERE project_id = $1;

-- Assign a member to a project
-- name: CreateProjectMember :exec
INSERT INTO project_member (member_id, project_id, role, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Get all members assigned to a project
-- name: GetProjectMembers :many
SELECT * FROM project_member WHERE project_id = $1;

-- Get a project's member by their ID
-- name: GetProjectMemberByID :one
SELECT * FROM project_member WHERE member_id = $1 AND project_id = $2;

-- Update a project's member information
-- name: UpdateProjectMemberByID :exec
UPDATE project_member SET role = $3, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1 AND project_id = $2;

-- Remove a member from a project
-- name: DeleteProjectMemberByID :exec
DELETE FROM project_member WHERE member_id = $1 AND project_id = $2;