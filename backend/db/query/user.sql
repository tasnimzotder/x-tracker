-- name: CreateUser :one
INSERT INTO
    users (
        username, hashed_password, email, created_at, role
    )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    email = $2,
    updated_at = $3,
    status = $4,
    role = $5,
    phone_number = $6,
    country_code = $7,
    first_name = $8,
    last_name = $9,
    postal_code = $10
WHERE
    id = $11 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT count(*) FROM users;