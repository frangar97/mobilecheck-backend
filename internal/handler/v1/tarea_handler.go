package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerTareasDelDiaMovil(ctx *gin.Context) {
	fecha := ctx.Query("fecha")
	usuarioId := ctx.GetInt64("usuarioId")

	tareas, err := h.services.TareaService.ObtenerTareasDelDia(ctx.Request.Context(), fecha, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tareas)
}

func (h *Handler) crearTareaMovil(ctx *gin.Context) {
	var tareaJSON model.CreateTareaModelMovil

	if err := ctx.BindJSON(&tareaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	usuarioId := ctx.GetInt64("usuarioId")
	nuevaTarea, err := h.services.TareaService.CrearTareaMovil(ctx.Request.Context(), tareaJSON, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevaTarea)
}

func (h *Handler) obtenerTareasWebCantidadPorUsuarioRangoFecha(ctx *gin.Context) {
	fechaInicio := ctx.Query("fechaInicio")
	fechaFin := ctx.Query("fechaFin")
	visitas, err := h.services.TareaService.ObtenerCantidadTareasUsuarioPorFecha(ctx.Request.Context(), fechaInicio, fechaFin)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, visitas)
}
