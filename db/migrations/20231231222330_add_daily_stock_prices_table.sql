-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE daily_stock_prices (
  id SERIAL PRIMARY KEY,
  stock_id INT NOT NULL,
  date DATE NOT NULL,
  open_price DECIMAL(10, 2) NOT NULL,
  high_price DECIMAL(10, 2) NOT NULL,
  low_price DECIMAL(10, 2) NOT NULL,
  close_price DECIMAL(10, 2) NOT NULL,
  volume BIGINT NOT NULL,
  FOREIGN KEY (stock_id) REFERENCES stocks (id) ON DELETE CASCADE
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE daily_stock_prices;
-- +goose StatementEnd
