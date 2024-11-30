package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"medods/internal/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) Refresh(tokens models.Tokens) (models.Tokens, error) {
	return models.Tokens{}, nil

}
func (u *UserPostgres) SetRefreshToken(token models.RefreshToken) error {
	query := `INSERT INTO users(id, refresh_token, exp,ip,email) 
				VALUES ($1,$2,$3,$4,$5)
				ON CONFLICT (id)
				DO UPDATE SET refresh_token = $6, exp = $7, ip = $8,email = $9;
 			`
	_, err := u.db.Exec(query, token.ID, token.Refresh_token, token.Exp, token.Ip, token.Email, token.Refresh_token, token.Exp, token.Ip, token.Email)
	return err
}
func (u *UserPostgres) GetRefreshToken(email string) (models.RefreshToken, error) {
	token := []models.RefreshToken{}
	query := `SELECT * FROM users WHERE email = $1`
	err := u.db.Select(&token, query, email)
	if err != nil {
		return models.RefreshToken{}, err
	} else if len(token) == 0 {
		return models.RefreshToken{}, fmt.Errorf("Токен не найден")
	}
	return token[0], nil

}
