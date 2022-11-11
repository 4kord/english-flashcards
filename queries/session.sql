-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1;

-- name: GetSessionByToken :one
SELECT * FROM sessions
WHERE refresh_token = $1;

-- name: CreateSession :one
INSERT INTO sessions (refresh_token, user_agent, user_id, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CountSessions :one
SELECT COUNT(*) FROM sessions
WHERE user_id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE refresh_token = $1;

-- name: DeleteOldestSession :exec
DELETE FROM sessions
WHERE created_at = (SELECT MIN(created_at) FROM sessions);