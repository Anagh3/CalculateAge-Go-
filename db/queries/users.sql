-- name: CreateUser :execresult
INSERT INTO users (name,dob) VALUES (?,?);

-- name: GetUserByID :one
SELECT id, name, dob FROM users WHERE id = ?;

-- name: GetAllUsers :many
SELECT id, name, dob FROM users;

-- name: UpdateUser :execresult
UPDATE users SET name = ?, dob = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;