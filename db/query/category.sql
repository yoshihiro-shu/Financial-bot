-- name: CreateCategory :one
INSERT INTO categories (
    name
) VALUES (
  $1
)
RETURNING *;

-- name: ListCategory :many
SELECT * FROM categories
ORDER BY id;
