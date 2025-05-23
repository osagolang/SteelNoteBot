-- +migrate Up
CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    telegram_id INTEGER NOT NULL REFERENCES users(telegram_id),
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    weight NUMERIC,
    reps INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_records_user ON records(telegram_id);
CREATE INDEX idx_records_exercise ON records(exercise_id);
CREATE INDEX idx_records_created ON records(created_at);

-- +migrate Down
DROP INDEX IF EXISTS idx_records_user;
DROP INDEX IF EXISTS idx_records_exercise;
DROP INDEX IF EXISTS idx_records_created;
DROP TABLE IF EXISTS records;
