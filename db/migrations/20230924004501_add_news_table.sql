-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255),
    link text NOT NULL UNIQUE,
    thumbnail text,
    score int NOT NULL DEFAULT 0,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    udpated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE news;
-- +goose StatementEnd
