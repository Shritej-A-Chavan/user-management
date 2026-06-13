-- name: CreateUser :execresult
INSERT INTO users (name, dob) VALUES (?, ?);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsers :many
SELECT * FROM users ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateUser :exec
UPDATE users SET name = ?, dob = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;