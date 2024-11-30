package service

import (
	"fmt"
	"log"
	"medods/internal/models"
	"medods/internal/repository"
	"medods/pkg/utils"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetAccessToken(id int, ip string) (string, error) {
	token, _, err := utils.GenerateToken(id, ip)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (u *UserService) GetRefreshToken(id int, ip string, email string) (string, error) {
	token := utils.GenerateRefreshToken()
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()
	hashToken, err := utils.HashToken(token)
	if err != nil {
		return "", err
	}
	refreshToken := models.RefreshToken{
		ID:            id,
		Refresh_token: hashToken,
		Ip:            ip,
		Exp:           exp,
		Email:         email,
	}
	err = u.repo.SetRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) Refresh(token string, ip string, email string) (models.Tokens, error) {

	refreshToken, err := u.repo.GetRefreshToken(email)
	if err != nil {
		return models.Tokens{}, err
	}
	if !utils.CompareHash(token, refreshToken.Refresh_token) {
		return models.Tokens{}, fmt.Errorf("Неправильный refresh токен")
	} else if refreshToken.Exp < time.Now().Unix() {
		return models.Tokens{}, fmt.Errorf("Срок действия токена истек")
	}
	if refreshToken.Ip == ip {
		SendEmail(email)
		return models.Tokens{}, fmt.Errorf("Странный ip-адрес")
	}
	access, err := u.GetAccessToken(refreshToken.ID, ip)
	if err != nil {
		return models.Tokens{}, err
	}
	refresh, err := u.GetRefreshToken(refreshToken.ID, ip, email)
	if err != nil {
		return models.Tokens{}, err
	}
	return models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
		Email:        email,
	}, nil
}

func SendEmail(email string) {
	log.Printf("ip-адреса пользователя %s не совпадают", email)
}
