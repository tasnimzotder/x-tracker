-- name: CreateAccess :one
INSERT INTO device_access (
    device_id, user_id, permission
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccessWithDeviceIDAndUserID :one
SELECT * FROM device_access WHERE device_id = $1 AND user_id = $2;

-- name: GetAccessWithDeviceID :many
SELECT * FROM device_access WHERE device_id = $1;

-- name: GetAccessWithUserID :many
SELECT * FROM device_access WHERE user_id = $1;

-- name: DeleteAccessWithDeviceID :many
DELETE FROM device_access WHERE device_id = $1
RETURNING *;

-- name: DeleteAccessWithUserID :many
DELETE FROM device_access WHERE user_id = $1
RETURNING *;