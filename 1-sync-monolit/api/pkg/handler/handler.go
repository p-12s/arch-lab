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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/user", h.userIdentity)
	{
		users.GET("/", h.getMe)
	}

	loyalty := router.Group("/loyalty", h.userIdentity)
	{
		loyalty.POST("/create_user_card", h.createCard)
	}

	notification := router.Group("/notification", h.userIdentity)
	{
		notification.POST("/send_email", h.sendEmail)
	}

	return router
}
