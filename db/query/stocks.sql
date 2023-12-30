-- name: CreateStocks :exec
INSERT INTO stocks (
    symbol, name, open, close, high, low, volume, date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
);
