-- name: CreateUser :one
INSERT INTO users (username, hashed_password, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :one
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
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;