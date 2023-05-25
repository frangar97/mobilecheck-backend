package v1

import (
	"net/http"
	"strconv"
	"strings"

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
	fecha := ctx.Query("fechaTarea") + " " + ctx.Query("horaTarea")
	responsable := ctx.Query("usuarioId")
	usuarioId, err := strconv.ParseInt(responsable, 0, 0)

	tareas, err := h.services.TareaService.VerificarTarea(ctx.Request.Context(), fecha, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tareas)
}

func (h *Handler) validarTareaExcel(ctx *gin.Context) {
	fecha := ctx.Query("fechaTarea") + " " + ctx.Query("horaTarea")

	responsable := ctx.Query("usuarioId")
	usuarioId, err := strconv.ParseInt(responsable, 0, 0)

	cliente := ctx.Query("clienteId")
	clienteId, err := strconv.ParseInt(cliente, 0, 0)

	tipoVisita := ctx.Query("tipoVisitaId")
	tipoVisitaId, err := strconv.ParseInt(tipoVisita, 0, 0)

	tareas, err := h.services.TareaService.ValidarDataExcel(ctx, clienteId, usuarioId, tipoVisitaId, fecha)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tareas)
}

func (h *Handler) obtenerTareasHorasWeb(ctx *gin.Context) {

	var tareaJSON model.ParamReportTareasHoras

	fechas := strings.Split(ctx.Query("fechas"), ",")
	responsables := strings.Split(ctx.Query("usuarioId"), ",")
	var responsableId = []int{}

	for _, i := range responsables {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		responsableId = append(responsableId, j)
	}

	tareaJSON.Fechas = fechas
	tareaJSON.UsuarioId = responsableId

	tareas, err := h.services.TareaService.ObtenerTareasHorasWeb(ctx, tareaJSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tareas)
}

func (h *Handler) eliminarTareas(ctx *gin.Context) {

	tareas := strings.Split(ctx.Query("tareas"), ",")
	var tareasId = []int64{}

	for _, i := range tareas {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		tareasId = append(tareasId, int64(j))
	}

	rows, err := h.services.TareaService.EliminarTareas(ctx, tareasId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, rows)
}
