package handler

import (
	"github.com/gin-gonic/gin"
	"medods/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetTokens(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "id должно быть числом")
		return
	}
	ip := c.ClientIP()

	accessToken, err := h.service.User.GetAccessToken(id, ip)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}
	refreshToken, err := h.service.User.GetRefreshToken(id, ip)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}
	c.AbortWithStatusJSON(http.StatusOK, tokens)
}

func (h *Handler) Refresh(c *gin.Context) {
	var tokens models.Tokens
	if err := c.ShouldBind(&tokens); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Не тот формат переданных параметров")
		return
	}
	tokens, err := h.service.Refresh(tokens)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "Ошибка обновления токена")
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, tokens)

}
