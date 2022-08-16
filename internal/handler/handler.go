package handler

import (
	v1 "github.com/frangar97/mobilecheck-backend/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *gin.Engine {
	mux := gin.Default()

	h.initAPI(mux)

	return mux
}

func (h *Handler) initAPI(group *gin.Engine) {
	handlerV1 := v1.NewHandler()

	api := group.Group("/api")
	{
		handlerV1.Init(api)
	}
}
