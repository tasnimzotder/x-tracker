// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"
)

type Device struct {
	ID         int64     `json:"id"`
	DeviceName string    `json:"device_name"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}

type DeviceAccess struct {
	ID          int64     `json:"id"`
	DeviceID    int64     `json:"device_id"`
	UserID      int64     `json:"user_id"`
	Permission  string    `json:"permission"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

type DeviceActivity struct {
	ID        int64     `json:"id"`
	DeviceID  int64     `json:"device_id"`
	CreatedAt time.Time `json:"created_at"`
	Panic     bool      `json:"panic"`
	Fall      bool      `json:"fall"`
}

type User struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
}
