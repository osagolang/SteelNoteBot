package services

import (
	"context"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type ExerciseRepo interface {
	GetExerciseByGroup(ctx context.Context, muscleGroup string) ([]models.Exercise, error)
}

type ExerciseService struct {
	Repo ExerciseRepo
}

func NewExerciseService(repo ExerciseRepo) *ExerciseService {
	return &ExerciseService{Repo: repo}
}

func (s *ExerciseService) GetExerciseByGroup(ctx context.Context, muscleGroup string) ([]models.Exercise, error) {
	return s.Repo.GetExerciseByGroup(ctx, muscleGroup)
}
