package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initMovilRoutes(c *gin.RouterGroup) {
	movil := c.Group("/movil")
	{
		movil.POST("/login", h.LoginWeb)

		tipoVisita := movil.Group("/tipovisita")
		{
			tipoVisita.GET("", h.obtenerTiposVisita)
		}
	}
}

func (h *Handler) LoginMovil(ctx *gin.Context) {
	var credenciales model.AuthCredencialModel

	if err := ctx.BindJSON(&credenciales); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.services.AuthService.LoginMovil(ctx.Request.Context(), credenciales)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
