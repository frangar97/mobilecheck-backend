package v1

import (
	"github.com/frangar97/mobilecheck-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		h.initWebRoutes(v1)
		h.initMovilRoutes(v1)
	}
}
