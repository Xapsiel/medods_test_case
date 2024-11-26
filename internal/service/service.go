package service

import (
	"medods/internal/models"
	"medods/internal/repository"
)

type User interface {
	GetAccessToken(id int, ip string) (string, error)
	GetRefreshToken(id int, ip string) (string, error)
	Refresh(tokens models.Tokens) (models.Tokens, error)
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{User: NewUserService(repo.User)}
}
