package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerCargoUsuario(ctx *gin.Context) {
	paises, err := h.services.CargoService.ObtenerCargoUsuario(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, paises)
}
