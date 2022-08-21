package handler

import (
	v1 "github.com/frangar97/mobilecheck-backend/internal/handler/v1"
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

func (h *Handler) Init() *gin.Engine {
	mux := gin.Default()

	mux.Use(corsMiddleware)

	h.initAPI(mux)

	return mux
}

func (h *Handler) initAPI(group *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)

	api := group.Group("/api")
	{
		handlerV1.Init(api)
	}
}
