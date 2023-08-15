package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerPaises(ctx *gin.Context) {
	paises, err := h.services.PaisService.ObtenerPaises(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, paises)
}
