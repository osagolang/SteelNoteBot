package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type RecordRepo struct {
	db *pgxpool.Pool
}

func NewRecordRepo(db *pgxpool.Pool) *RecordRepo {
	return &RecordRepo{db: db}
}

func (r *RecordRepo) GetRecords(ctx context.Context, userID int64, exerciseID int, limit int) ([]models.Record, error) {

	rows, err := r.db.Query(ctx, `
		SELECT user_id, exercise_id, weight, reps, created_at
		FROM records
		WHERE user_id = $1 AND exercise_id = $2
		ORDER BY created_at DESC
		LIMIT $3`, userID, exerciseID, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var rec models.Record
		if err := rows.Scan(&rec.UserID, &rec.ExerciseID, &rec.Weight, &rec.Reps, &rec.CreatedAt); err != nil {
			return nil, err
		}

		records = append(records, rec)
	}

	return records, nil
}

func (r *RecordRepo) SaveRecord(ctx context.Context, rec models.Record) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO records(user_id, exercise_id, weight, reps, created_at)
		VALUES ($1, $2, $3, $4, $5)
		`, rec.UserID, rec.ExerciseID, rec.Weight, rec.Reps, rec.CreatedAt)

	return err
}
