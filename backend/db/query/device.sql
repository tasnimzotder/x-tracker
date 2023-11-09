-- name: CreateDevice :one
INSERT INTO devices (
    device_name, created_at, status
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetDevice :one
SELECT * FROM devices WHERE id = $1;

-- name: UpdateDeviceStatus :one
UPDATE devices SET status = $1 WHERE id = $2 RETURNING *;