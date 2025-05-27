package repositories

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type RecordRepo struct {
	db *pgxpool.Pool
}

func NewRecordRepo(db *pgxpool.Pool) *RecordRepo {
	return &RecordRepo{db: db}
}

func (r *RecordRepo) GetRecords(ctx context.Context, telegramID int64, exerciseID int, limit int) ([]models.Record, error) {

	rows, err := r.db.Query(ctx, `
		SELECT telegram_id, exercise_id, weight, reps, created_at, e.name, e.muscle_group, e.has_weight
		FROM records
		INNER JOIN exercises e ON e.id = records.exercise_id
		WHERE telegram_id = $1 AND exercise_id = $2
		ORDER BY created_at DESC
		LIMIT $3`, telegramID, exerciseID, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var rec models.Record
		var ex models.Exercise
		if err := rows.Scan(&rec.TelegramID, &ex.ID, &rec.Weight, &rec.Reps, &rec.CreatedAt, &ex.Name, &ex.MuscleGroup, &ex.HasWeight); err != nil {
			return nil, err
		}
		rec.Exercise = ex

		records = append(records, rec)
	}

	return records, nil
}

func (r *RecordRepo) SaveRecord(ctx context.Context, rec models.Record) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO records(telegram_id, exercise_id, weight, reps, created_at)
		VALUES ($1, $2, $3, $4, $5)
		`, rec.TelegramID, rec.Exercise.ID, rec.Weight, rec.Reps, rec.CreatedAt)

	return err
}

func (r *RecordRepo) GetBestRecords(ctx context.Context, telegramID int64, exerciseID int) (*models.Record, error) {

	var rec models.Record
	var ex models.Exercise

	err := r.db.QueryRow(ctx, `
		SELECT weight, reps, created_at, e.id, e.name, e.muscle_group, e.has_weight
		FROM records
		INNER JOIN exercises e ON e.id = records.exercise_id
		WHERE telegram_id = $1 AND exercise_id = $2
		ORDER BY weight DESC, created_at DESC
		LIMIT 1`,
		telegramID, exerciseID,
	).Scan(&rec.Weight, &rec.Reps, &rec.CreatedAt, &ex.ID, &ex.Name, &ex.MuscleGroup, &ex.HasWeight)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	rec.Exercise = ex

	return &rec, nil
}
