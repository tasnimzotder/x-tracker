// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: device_access.sql

package db

import (
	"context"
)

const createAccess = `-- name: CreateAccess :one
INSERT INTO device_access (
    device_id, user_id, permission
) VALUES (
    $1, $2, $3
) RETURNING id, device_id, user_id, permission, created_at, last_updated
`

type CreateAccessParams struct {
	DeviceID   int64  `json:"device_id"`
	UserID     int64  `json:"user_id"`
	Permission string `json:"permission"`
}

func (q *Queries) CreateAccess(ctx context.Context, arg CreateAccessParams) (DeviceAccess, error) {
	row := q.db.QueryRow(ctx, createAccess, arg.DeviceID, arg.UserID, arg.Permission)
	var i DeviceAccess
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.UserID,
		&i.Permission,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const deleteAccessWithDeviceID = `-- name: DeleteAccessWithDeviceID :many
DELETE FROM device_access WHERE device_id = $1
RETURNING id, device_id, user_id, permission, created_at, last_updated
`

func (q *Queries) DeleteAccessWithDeviceID(ctx context.Context, deviceID int64) ([]DeviceAccess, error) {
	rows, err := q.db.Query(ctx, deleteAccessWithDeviceID, deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DeviceAccess{}
	for rows.Next() {
		var i DeviceAccess
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.UserID,
			&i.Permission,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteAccessWithUserID = `-- name: DeleteAccessWithUserID :many
DELETE FROM device_access WHERE user_id = $1
RETURNING id, device_id, user_id, permission, created_at, last_updated
`

func (q *Queries) DeleteAccessWithUserID(ctx context.Context, userID int64) ([]DeviceAccess, error) {
	rows, err := q.db.Query(ctx, deleteAccessWithUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DeviceAccess{}
	for rows.Next() {
		var i DeviceAccess
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.UserID,
			&i.Permission,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccessWithDeviceID = `-- name: GetAccessWithDeviceID :many
SELECT id, device_id, user_id, permission, created_at, last_updated FROM device_access WHERE device_id = $1
`

func (q *Queries) GetAccessWithDeviceID(ctx context.Context, deviceID int64) ([]DeviceAccess, error) {
	rows, err := q.db.Query(ctx, getAccessWithDeviceID, deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DeviceAccess{}
	for rows.Next() {
		var i DeviceAccess
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.UserID,
			&i.Permission,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccessWithDeviceIDAndUserID = `-- name: GetAccessWithDeviceIDAndUserID :many
SELECT id, device_id, user_id, permission, created_at, last_updated FROM device_access WHERE device_id = $1 AND user_id = $2
`

type GetAccessWithDeviceIDAndUserIDParams struct {
	DeviceID int64 `json:"device_id"`
	UserID   int64 `json:"user_id"`
}

func (q *Queries) GetAccessWithDeviceIDAndUserID(ctx context.Context, arg GetAccessWithDeviceIDAndUserIDParams) ([]DeviceAccess, error) {
	rows, err := q.db.Query(ctx, getAccessWithDeviceIDAndUserID, arg.DeviceID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DeviceAccess{}
	for rows.Next() {
		var i DeviceAccess
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.UserID,
			&i.Permission,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccessWithUserID = `-- name: GetAccessWithUserID :many
SELECT id, device_id, user_id, permission, created_at, last_updated FROM device_access WHERE user_id = $1
`

func (q *Queries) GetAccessWithUserID(ctx context.Context, userID int64) ([]DeviceAccess, error) {
	rows, err := q.db.Query(ctx, getAccessWithUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DeviceAccess{}
	for rows.Next() {
		var i DeviceAccess
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.UserID,
			&i.Permission,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
