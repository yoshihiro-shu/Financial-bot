-- name: CreateProvider :one
INSERT INTO providers (
    name
) VALUES (
  $1
)
RETURNING *;

-- name: ListProvider :many
SELECT * FROM providers
ORDER BY id;
