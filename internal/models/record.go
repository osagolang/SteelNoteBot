package models

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DB = *pgxpool.Pool

type Record struct {
	TelegramID int64
	Exercise   Exercise
	Weight     *float64
	Reps       int
	CreatedAt  time.Time
}

func (r Record) FormatLastMsg() string {
	if r.Weight == nil {
		return fmt.Sprintf("%s - %d раз(а)\n", r.CreatedAt.Format("02.01"), r.Reps)
	}

	return fmt.Sprintf("%s - %.1f кг. x %d раз(а)\n", r.CreatedAt.Format("02.01"), *r.Weight, r.Reps)
}

func (r Record) FormatMsg() string {
	if r.Weight == nil {
		return fmt.Sprintf("%s: %d раз\n", r.Exercise.Name, r.Reps)
	}

	return fmt.Sprintf("%s: %.1f кг × %d раз\n", r.Exercise.Name, *r.Weight, r.Reps)
}
