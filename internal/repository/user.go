package repository

import (
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
	query := `INSERT INTO users(id, refresh_token, exp,ip) 
				VALUES ($1,$2,$3,$4)
				ON CONFLICT 
				DO UPDATE SET refresh_token = $5, exp = $6, ip = $7;
 			`
	_, err := u.db.Exec(query, token.ID, token.Refresh_token, token.Exp, token.Ip, token.Refresh_token, token.Refresh_token, token.Ip)
	return err
}
func (u *UserPostgres) GetRefreshToken(id int) (models.RefreshToken, error) {
	token := models.RefreshToken{}
	query := `SELECT refresh_token, exp,ip FROM users WHERE id = '$1'`
	err := u.db.Select(&token, query, id)
	if err != nil {
		return token, err
	}
	return token, nil

}
