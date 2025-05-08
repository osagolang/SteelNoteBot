-- +migrate Up
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    muscle_group TEXT NOT NULL,
    has_weight BOOLEAN NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS exercises;
