-- name: CreateProvider :one
INSERT INTO providers (
    name
) VALUES (
  $1
)
RETURNING *;