package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type ExerciseRepo struct {
	db *pgxpool.Pool
}

func NewExerciseRepo(db *pgxpool.Pool) *ExerciseRepo {
	return &ExerciseRepo{db: db}
}

func (r *ExerciseRepo) GetExerciseByGroup(ctx context.Context, muscleGroup string) ([]models.Exercise, error) {

	rows, err := r.db.Query(ctx,
		`SELECT id, name FROM exercises WHERE muscle_group = $1 ORDER BY id`, muscleGroup)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var ex models.Exercise
		if err := rows.Scan(&ex.ID, &ex.Name); err != nil {
			return nil, err
		}
		exercises = append(exercises, ex)
	}
	return exercises, nil
}
