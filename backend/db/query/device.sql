-- name: CreateDevice :one
INSERT INTO
    devices (
        device_name, device_key, user_id, created_at, status
    )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetDevice :one
SELECT * FROM devices WHERE id = $1;

-- name: GetDevicesByUser :many
SELECT * FROM devices WHERE user_id = $1;