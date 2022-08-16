package v1

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		h.initUsuarioRoutes(v1)
	}
}
