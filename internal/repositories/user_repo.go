package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (telegram_id, username) 
		 VALUES ($1, $2) 
		 ON CONFLICT (telegram_id) DO UPDATE SET last_time = NOW()`,
		user.TelegramID, user.Username,
	)
	return err
}
