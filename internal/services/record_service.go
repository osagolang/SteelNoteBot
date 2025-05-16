package services

import (
	"context"
	"github.com/osagolang/SteelNoteBot/internal/models"
	"time"
)

type RecordRepo interface {
	GetRecords(ctx context.Context, userID int64, exerciseID int, limit int) ([]models.Record, error)
	SaveRecord(ctx context.Context, record models.Record) error
}

type RecordService struct {
	Repo RecordRepo
}

func NewRecordService(repo RecordRepo) *RecordService {
	return &RecordService{Repo: repo}
}

func (s *RecordService) GetRecords(ctx context.Context, userID int64, exerciseID int, limit int) ([]models.Record, error) {
	return s.Repo.GetRecords(ctx, userID, exerciseID, limit)
}

func (s *RecordService) AddRecord(ctx context.Context, userID int64, exerciseID int, weight float64, reps int) error {
	rec := models.Record{
		UserID:     userID,
		ExerciseID: exerciseID,
		Weight:     weight,
		Reps:       reps,
		CreatedAt:  time.Now(),
	}

	return s.Repo.SaveRecord(ctx, rec)
}
