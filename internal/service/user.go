package service

import (
	"medods/internal/models"
	"medods/internal/repository"
	"medods/pkg/utils"
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
func (u *UserService) GetRefreshToken(id int, ip string) (string, error) {
	token, time, err := utils.GenerateToken(id, ip)
	if err != nil {
		return "", err
	}
	hashToken, err := utils.HashToken(token)
	if err != nil {
		return "", err
	}
	refreshToken := models.RefreshToken{
		ID:            id,
		Refresh_token: hashToken,
		Ip:            ip,
		Exp:           time,
	}
	err = u.repo.SetRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) Refresh(tokens models.Tokens) (models.Tokens, error) {
	access := tokens.AccessToken
	refresh := tokens.RefreshToken
	a_user_id, a_ip, err := utils.ExtractPayload(access)
	if err != nil {
		return models.Tokens{}, err
	}
	r_user_id, r_ip, err := utils.ExtractPayload(refresh)

	return u.repo.Refresh(tokens)
}
