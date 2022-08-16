package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsuarioRoutes(api *gin.RouterGroup) {
	users := api.Group("/usuario")
	{
		users.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Ruta usuario"})
		})
	}
}
