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

func (h *Handler) crearTareaWeb(ctx *gin.Context) {
	var tareaJSON model.CreateTareaModelWeb

	if err := ctx.BindJSON(&tareaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	nuevaTarea, err := h.services.TareaService.CrearTareaWeb(ctx.Request.Context(), tareaJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevaTarea)
}

func (h *Handler) obtenerTareasWeb(ctx *gin.Context) {
	fechaInicio := ctx.Query("fechaInicio")
	fechaFin := ctx.Query("fechaFin")
	visitas, err := h.services.TareaService.ObtenerTareasWeb(ctx.Request.Context(), fechaInicio, fechaFin)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, visitas)
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

func (h *Handler) completarTarea(ctx *gin.Context) {
	var tareaJSON model.CompletarTareaModel

	if err := ctx.Bind(&tareaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	usuarioId := ctx.GetInt64("usuarioId")

	err := h.services.TareaService.CompletarTarea(ctx, tareaJSON, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "creado con exito"})
}

func (h *Handler) crearTareaMasivaWeb(ctx *gin.Context) {
	var tareaJSON model.CreateTareaMasivaModelWeb

	if err := ctx.BindJSON(&tareaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	err := h.services.TareaService.CrearTareaMasivaWeb(ctx.Request.Context(), tareaJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Tareas creadas con exito"})
}

func (h *Handler) crearTareaMasivaExcelWeb(ctx *gin.Context) {
	var tareaJSON model.CreateTareasExcelWeb

	if err := ctx.BindJSON(&tareaJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	err := h.services.TareaService.CrearTareaMasivaExcelWeb(ctx.Request.Context(), tareaJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Tareas creadas con exito"})
}

func (h *Handler) verificarTarea(ctx *gin.Context) {
	// fechaTarea := ctx.Query("fechaTarea")
	fecha := ctx.Query("fechaTarea") + " " + ctx.Query("horaTarea")
	usuarioId := ctx.GetInt64("usuarioId")

	tareas, err := h.services.TareaService.VerificarTarea(ctx.Request.Context(), fecha, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tareas)
}
