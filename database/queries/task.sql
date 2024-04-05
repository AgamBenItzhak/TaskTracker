-- Insert a new Task Group into the database
-- name: CreateTaskGroup :one
INSERT INTO task_group (project_id, group_name, description, status, priority, start_date, end_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING task_group_id;

-- Get all Task Group IDs for a project
-- name: GetAllTaskGroupIDs :many
SELECT task_group_id FROM task_group WHERE project_id = $1;

-- Get all Task Groups for a project
-- name: GetAllTaskGroups :many
SELECT * FROM task_group WHERE project_id = $1;

-- Get a Task Group by its ID
-- name: GetTaskGroupByID :one
SELECT * FROM task_group WHERE task_group_id = $1;

-- Update a Task Group's information
-- name: UpdateTaskGroupByID :exec
UPDATE task_group SET group_name = $2, description = $3, status = $4, priority = $5, start_date = $6, end_date = $7, updated_at = CURRENT_TIMESTAMP
WHERE task_group_id = $1;

-- Delete a Task Group from the database
-- name: DeleteTaskGroupByID :exec
DELETE FROM task_group WHERE task_group_id = $1;

-- Insert a new member into a Task Group
-- name: CreateTaskGroupMember :exec
INSERT INTO task_group_member (member_id, task_group_id, role, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Get all member IDs assigned to a Task Group
-- name: GetAllTaskGroupMemberIDs :many
SELECT member_id FROM task_group_member WHERE task_group_id = $1;

-- Get all members assigned to a Task Group
-- name: GetAllTaskGroupMembers :many
SELECT * FROM task_group_member WHERE task_group_id = $1;

-- Get a member's Task Group by their ID
-- name: GetTaskGroupMemberByID :one
SELECT * FROM task_group_member WHERE member_id = $1 AND task_group_id = $2;

-- Update a member's Task Group information
-- name: UpdateTaskGroupMemberByID :exec
UPDATE task_group_member SET role = $3, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1 AND task_group_id = $2;

-- Remove a member from a Task Group
-- name: DeleteTaskGroupMemberByID :exec
DELETE FROM task_group_member WHERE member_id = $1 AND task_group_id = $2;

-- Insert a new Task into the database
-- name: CreateTask :one
INSERT INTO task (task_group_id, task_name, description, status, priority, start_date, end_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING task_id;

-- Get all Task IDs for a Task Group
-- name: GetAllTaskIDs :many
SELECT task_id FROM task WHERE task_group_id = $1;

-- Get all Tasks for a Task Group
-- name: GetAllTasks :many
SELECT * FROM task WHERE task_group_id = $1;

-- Get a Task by its ID
-- name: GetTaskByID :one
SELECT * FROM task WHERE task_id = $1;

-- Update a Task's information
-- name: UpdateTaskByID :exec
UPDATE task SET task_name = $2, description = $3, status = $4, priority = $5, start_date = $6, end_date = $7, updated_at = CURRENT_TIMESTAMP
WHERE task_id = $1;

-- Delete a Task from the database
-- name: DeleteTaskByID :exec
DELETE FROM task WHERE task_id = $1;

-- Insert a new member into a Task
-- name: CreateTaskMember :exec
INSERT INTO task_member (member_id, task_id, role, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Get all member IDs assigned to a Task
-- name: GetAllTaskMemberIDs :many
SELECT member_id FROM task_member WHERE task_id = $1;

-- Get all members assigned to a Task
-- name: GetAllTaskMembers :many
SELECT * FROM task_member WHERE task_id = $1;

-- Get a member's Task by their ID
-- name: GetTaskMemberByID :one
SELECT * FROM task_member WHERE member_id = $1 AND task_id = $2;

-- Update a member's Task information
-- name: UpdateTaskMemberByID :exec
UPDATE task_member SET role = $3, updated_at = CURRENT_TIMESTAMP
WHERE member_id = $1 AND task_id = $2;

-- Remove a member from a Task
-- name: DeleteTaskMemberByID :exec
DELETE FROM task_member WHERE member_id = $1 AND task_id = $2;