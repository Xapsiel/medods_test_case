package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"medods/internal/models"
	"net/http"
	"strconv"
)

// @Summary		Получение access и refresh токена
// @Tags			auth
// @Description Обновление access токена
// @Accept			json
// @Produce		json
// @Param		id	query		string	false	"ID пользователя"	default(1)
// @Param		email	query		string	false	"email пользователя"	default(DefaultEmail@mail.ru)
// @Success		200		{object}	models.Tokens
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/auth/getTokens [get]
func (h *Handler) GetTokens(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	email := c.DefaultQuery("email", "DefaultEmail@mail.ru")
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "id должно быть числом")
		return
	}
	ip := c.ClientIP()

	accessToken, err := h.service.User.GetAccessToken(id, ip)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	refreshToken, err := h.service.User.GetRefreshToken(id, ip, email)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	refreshToken = base64.StdEncoding.EncodeToString([]byte(refreshToken))
	Token := models.Tokens{RefreshToken: refreshToken, AccessToken: accessToken, Email: email}

	c.AbortWithStatusJSON(http.StatusOK, Token)
}

// @Summary Refresh Tokens
// @Description Refresh the access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param token body  models.Tokens true "Рефреш токен и email"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} handler.errorResponse
// @Failure 500 {object} handler.errorResponse
// @Router /auth/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
	var tokens models.Tokens
	if err := c.ShouldBind(&tokens); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Не тот формат переданных параметров")
		return
	}
	ip := c.ClientIP()
	refreshToken, err := base64.StdEncoding.DecodeString(tokens.RefreshToken)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "Ошибка конвертирования токена")
		return
	}
	tokens, err = h.service.Refresh(string(refreshToken), ip, tokens.Email)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "Ошибка обновления токена")
		return
	}
	tokens.RefreshToken = base64.StdEncoding.EncodeToString([]byte(tokens.RefreshToken))
	c.AbortWithStatusJSON(http.StatusOK, tokens)

}
