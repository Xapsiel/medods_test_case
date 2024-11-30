package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "medods/docs"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		auth.GET("/getTokens", h.GetTokens)
		auth.POST("/refresh", h.Refresh)
	}
	return router

}
