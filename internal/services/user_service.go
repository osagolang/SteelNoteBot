package services

import (
	"context"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type UserService struct {
	Repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) error {
	return s.Repo.CreateUser(ctx, user)
}
