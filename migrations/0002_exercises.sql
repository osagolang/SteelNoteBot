-- +migrate Up
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    muscle_group VARCHAR(255) NOT NULL,
    has_weight BOOLEAN NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS exercises;
