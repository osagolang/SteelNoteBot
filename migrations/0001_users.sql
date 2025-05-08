-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE IF EXISTS users;
