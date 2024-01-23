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
    hashed_password = $2,
    email = $3,
    updated_at = $4,
    status = $5,
    role = $6,
    phone_number = $7,
    country_code = $8,
    first_name = $9,
    last_name = $10,
    postal_code = $11
WHERE
    id = $12 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT count(*) FROM users;