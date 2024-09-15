-- name: GetUserByID :one
SELECT user_id, mail, name, hashed_password FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByIDWithEisaFiles :one
SELECT u.user_id, u.mail, u.name, u.hashed_password, ef.file_path
FROM users u
INNER JOIN eisa_files ef USING (user_id)
WHERE u.user_id = $1;

-- name: CreateUser :one
INSERT INTO users (
  user_id, mail, name, hashed_password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set mail = $2,
  name = $3,
  hashed_password = $4
WHERE user_id = $1;

-- name: UpsertEisaFile :exec
WITH updated AS (
  UPDATE eisa_files
  SET deleted_at = CURRENT_TIMESTAMP
  WHERE user_id = $1
  RETURNING *
)
INSERT INTO eisa_files (user_id, file_path)
VALUES ($1, $2);

-- name: UpdatePassword :exec
UPDATE users
  set hashed_password = $2
WHERE user_id = $1;

