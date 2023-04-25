package v1

import (
	"net/http"
	"strconv"

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

	visita, err := h.services.VisitaService.CrearVisita(ctx, visitaJSON, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, visita)
}

func (h *Handler) obtenerVisitasMovil(ctx *gin.Context) {
	usuarioId := ctx.GetInt64("usuarioId")
	fecha := ctx.Query("fecha")
	visitas, err := h.services.VisitaService.ObtenerVisitasPorUsuarioDelDia(ctx.Request.Context(), fecha, usuarioId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}

func (h *Handler) obtenerVisitasWebRangoFecha(ctx *gin.Context) {
	fechaInicio := ctx.Query("fechaInicio")
	fechaFin := ctx.Query("fechaFin")
	visitas, err := h.services.VisitaService.ObtenerVisitasPorRangoFecha(ctx.Request.Context(), fechaInicio, fechaFin)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}

func (h *Handler) obtenerVisitasWebCantidadPorUsuarioRangoFecha(ctx *gin.Context) {
	fechaInicio := ctx.Query("fechaInicio")
	fechaFin := ctx.Query("fechaFin")
	visitas, err := h.services.VisitaService.ObtenerCantidadVisitaPorUsuario(ctx.Request.Context(), fechaInicio, fechaFin)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}

func (h *Handler) obtenerVisitasWebCantidadPorTipoRangoFecha(ctx *gin.Context) {
	fechaInicio := ctx.Query("fechaInicio")
	fechaFin := ctx.Query("fechaFin")
	visitas, err := h.services.VisitaService.ObtenerCantidadVisitaPorTipo(ctx.Request.Context(), fechaInicio, fechaFin)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}

func (h *Handler) obtenerVisitaTareaWeb(ctx *gin.Context) {

	idTarea := ctx.Param("id")
	tarea, err := strconv.ParseInt(idTarea, 0, 0)

	visitas, err := h.services.VisitaService.ObtenerVisitaTarea(ctx.Request.Context(), tarea)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}
