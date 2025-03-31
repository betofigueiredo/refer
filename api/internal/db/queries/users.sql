-- name: CreateUser :one
INSERT INTO
  users (id, name, email)
VALUES
  (?, ?, ?) RETURNING *;

-- name: GetUser :one
SELECT
  *
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: GetUsers :many
SELECT
  *
FROM
  users
ORDER BY
  created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET
  name = ?,
  updated_at = NOW ()
WHERE
  id = ? RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
  id = ?;
