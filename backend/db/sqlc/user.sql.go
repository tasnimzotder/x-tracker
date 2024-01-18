// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING id, username, hashed_password, email, created_at, last_updated_at, phone_number, country_code, first_name, last_name, postal_code
`

type CreateUserParams struct {
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.PhoneNumber,
		&i.CountryCode,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, hashed_password, email, created_at, last_updated_at, phone_number, country_code, first_name, last_name, postal_code FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.PhoneNumber,
		&i.CountryCode,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, hashed_password, email, created_at, last_updated_at, phone_number, country_code, first_name, last_name, postal_code FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.PhoneNumber,
		&i.CountryCode,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, hashed_password, email, created_at, last_updated_at, phone_number, country_code, first_name, last_name, postal_code FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.PhoneNumber,
		&i.CountryCode,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    hashed_password = $2,
    email = $3,
    last_updated_at = $4,
    phone_number = $5,
    country_code = $6,
    first_name = $7,
    last_name = $8,
    postal_code = $9
WHERE id = $7
RETURNING id, username, hashed_password, email, created_at, last_updated_at, phone_number, country_code, first_name, last_name, postal_code
`

type UpdateUserParams struct {
	Username       string      `json:"username"`
	HashedPassword string      `json:"hashed_password"`
	Email          string      `json:"email"`
	LastUpdatedAt  time.Time   `json:"last_updated_at"`
	PhoneNumber    pgtype.Int8 `json:"phone_number"`
	CountryCode    pgtype.Int4 `json:"country_code"`
	FirstName      pgtype.Text `json:"first_name"`
	LastName       pgtype.Text `json:"last_name"`
	PostalCode     pgtype.Int8 `json:"postal_code"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.LastUpdatedAt,
		arg.PhoneNumber,
		arg.CountryCode,
		arg.FirstName,
		arg.LastName,
		arg.PostalCode,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.PhoneNumber,
		&i.CountryCode,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
	)
	return i, err
}
