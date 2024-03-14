-- name: CreateGeofence :one
INSERT INTO 
    geofences (
    device_id, 
    geofence_name,
    status,
    rule,
    created_at,
    updated_at,
    center_lat,
    center_long,
    radius
    )
VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9
) 
RETURNING *; 

-- name: GetGeofencesByDevice :many 
SELECT * 
FROM geofences 
WHERE device_id = $1; 

-- name: GetGeofence :one
SELECT *
FROM geofences
WHERE id = $1;

-- name: UpdateGeofence :one
UPDATE geofences
SET 
    geofence_name = $2,
    status = $3,
    updated_at = $4,
    center_lat = $5,
    center_long = $6,
    radius = $7
WHERE id = $1
RETURNING *;

-- name: DeleteGeofence :one
DELETE FROM geofences
WHERE id = $1
RETURNING *;
