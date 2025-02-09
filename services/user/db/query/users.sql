-- name: CreateUser :one
INSERT INTO users(
    name,
    email,
    password,
    phone_number
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET
    password = COALESCE(sqlc.narg(password), password),
    name = COALESCE(sqlc.narg(name), name),
    email = COALESCE(sqlc.narg(email), email),
    phone_number = COALESCE(sqlc.narg(phone_number), phone_number)
WHERE id = $1
RETURNING *;