package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) actualizarSubsidioTelefoniaImpulsadoras(ctx *gin.Context) {
	result, err := h.services.ImportarExportarDataService.ActualizarDataImpulsadoras(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
