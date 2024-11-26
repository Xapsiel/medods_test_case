package handler

import (
	"github.com/gin-gonic/gin"
	"medods/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("auth")
	{
		auth.GET("GetTokens", h.GetTokens)
	}
	return router

}
