-- name: CreateDeviceActivity :one
INSERT INTO device_activities (
  device_id,
  panic,
  fall
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetDeviceActivityPanic :one
SELECT * FROM device_activities
WHERE device_id = $1 AND panic = true
ORDER BY created_at DESC
LIMIT 1;

-- name: GetDeviceActivity :many
SELECT * FROM device_activities
WHERE device_id = $1
ORDER BY created_at DESC
LIMIT $2;