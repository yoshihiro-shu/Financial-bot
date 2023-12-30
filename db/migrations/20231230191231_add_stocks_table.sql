-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    symbol varchar(5) NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    open FLOAT(2) NOT NULL,
    close FLOAT(2) NOT NULL,
    high FLOAT(2) NOT NULL,
    low FLOAT(2) NOT NULL,
    volume INT NOT NULL,
    date DATE NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE stocks;
-- +goose StatementEnd
