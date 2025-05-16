package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DB = *pgxpool.Pool

type Record struct {
	UserID     int64
	ExerciseID int
	Weight     float64
	Reps       int
	CreatedAt  time.Time
}
