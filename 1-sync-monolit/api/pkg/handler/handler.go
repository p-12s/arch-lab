package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.MaxMultipartMemory = 1 << 20 // 1 MiB (default is 32 MiB)

	router.GET("/health", h.health)
	router.POST("/api/registerUser", h.registerUser)

	api := router.Group("/api/v1", h.userIdentity)
	{
		api.GET("/files", h.health)
	}

	return router
}
