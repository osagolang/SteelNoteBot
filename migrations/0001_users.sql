-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    username VARCHAR(255) UNIQUE,
    first_time TIMESTAMP DEFAULT now(),
    last_time TIMESTAMP DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS users;
