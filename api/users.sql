-- name: CreateUser :one
INSERT INTO users (
  id,
  name,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
  name = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
