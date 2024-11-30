package repository

import (
	"github.com/jmoiron/sqlx"
	"medods/internal/models"
)

type User interface {
	Refresh(tokens models.Tokens) (models.Tokens, error)
	SetRefreshToken(token models.RefreshToken) error
	GetRefreshToken(email string) (models.RefreshToken, error)
}
type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
