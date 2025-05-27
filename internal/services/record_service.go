package services

import (
	"context"
	"github.com/osagolang/SteelNoteBot/internal/models"
	"time"
)

type RecordRepo interface {
	GetRecords(ctx context.Context, telegramID int64, exerciseID int, limit int) ([]models.Record, error)
	SaveRecord(ctx context.Context, record models.Record) error
	GetBestRecords(ctx context.Context, telegramID int64, exerciseID int) (*models.Record, error)
}

type RecordService struct {
	Repo RecordRepo
}

func NewRecordService(repo RecordRepo) *RecordService {
	return &RecordService{Repo: repo}
}

func (s *RecordService) GetRecords(ctx context.Context, telegramID int64, exerciseID int, limit int) ([]models.Record, error) {
	return s.Repo.GetRecords(ctx, telegramID, exerciseID, limit)
}

func (s *RecordService) AddRecord(ctx context.Context, telegramID int64, exerciseID int, weight *float64, reps int) error {
	rec := models.Record{
		TelegramID: telegramID,
		Exercise:   models.Exercise{ID: exerciseID},
		Weight:     weight,
		Reps:       reps,
		CreatedAt:  time.Now(),
	}

	return s.Repo.SaveRecord(ctx, rec)
}

func (s *RecordService) GetBestResult(ctx context.Context, telegramID int64, exerciseID int) (*models.Record, error) {
	return s.Repo.GetBestRecords(ctx, telegramID, exerciseID)
}
