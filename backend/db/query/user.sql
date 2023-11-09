-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;