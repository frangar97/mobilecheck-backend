package v1

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ObtenerImpulsadorasSubcidioTelefono(ctx *gin.Context) {

	data, err := h.services.ReporteService.ObtenerImpulsadorasSubcidioTelefono(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
