-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    symbol varchar(5) NOT NULL UNIQUE,
    name varchar(255) NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE stocks;
-- +goose StatementEnd
