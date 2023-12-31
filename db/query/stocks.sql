-- name: CreateStocks :exec
INSERT INTO stocks (
    symbol, name
) VALUES (
    $1, $2
);
