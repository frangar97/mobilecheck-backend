package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CrearVisita(ctx *gin.Context) {
	var visitaJSON model.CreateVisitaModel

	if err := ctx.Bind(&visitaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	usuarioId := ctx.GetInt64("usuarioId")

	idGenerado, err := h.services.VisitaService.CrearVisita(ctx, visitaJSON, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": idGenerado})
}
