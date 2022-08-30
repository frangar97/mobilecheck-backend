package v1

import (
	"net/http"
	"strconv"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerTiposVisita(ctx *gin.Context) {
	tiposVisita, err := h.services.TipoVisitaService.ObtenerTiposVisita(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tiposVisita)
}

func (h *Handler) obtenerTiposVisitaActiva(ctx *gin.Context) {
	tiposVisita, err := h.services.TipoVisitaService.ObtenerTiposVisitaActiva(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tiposVisita)
}

func (h *Handler) crearTipoVisita(ctx *gin.Context) {
	var tipoVisitaJSON model.CreateTipoVisitaModel

	if err := ctx.BindJSON(&tipoVisitaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	nuevoTipoVisita, err := h.services.TipoVisitaService.CrearTipoVisita(ctx.Request.Context(), tipoVisitaJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevoTipoVisita)

}

func (h *Handler) actualizarTipoVisita(ctx *gin.Context) {
	tipoVisitaIdParam := ctx.Param("tipoVisitaId")

	tipoVisitaId, err := strconv.ParseInt(tipoVisitaIdParam, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var tiposVisitaJSON model.UpdateTipoVisitaModel

	if err := ctx.BindJSON(&tiposVisitaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	actualizado, err := h.services.TipoVisitaService.ActualizarTipoVisita(ctx.Request.Context(), tipoVisitaId, tiposVisitaJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !actualizado {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo actualizar el tipo visita"})
		return
	}

	ctx.Status(http.StatusOK)
}
