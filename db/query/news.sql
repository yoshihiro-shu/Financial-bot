-- name: GetNews :one
SELECT * FROM news
WHERE id = $1 LIMIT 1;

-- name: ListNews :many
SELECT * FROM news
ORDER BY published_at;

-- name: CreateNews :one
INSERT INTO news (
  title, link, published_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteNews :exec
DELETE FROM news
WHERE id = $1;