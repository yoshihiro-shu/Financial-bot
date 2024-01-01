-- name: GetNews :one
SELECT * FROM news
WHERE id = $1 LIMIT 1;

-- name: GetNewsByTitle :one
SELECT 1 FROM news
WHERE title = $1 LIMIT 1;

-- name: ListNews :many
SELECT * FROM news
ORDER BY published_at;

-- name: CreateNews :one
INSERT INTO news (
  title, description, link, thumbnail, score, published_at, provider_id, category_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: DeleteNews :exec
DELETE FROM news
WHERE id = $1;