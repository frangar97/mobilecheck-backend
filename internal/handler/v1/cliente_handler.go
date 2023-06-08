package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerClientes(ctx *gin.Context) {
	clientes, err := h.services.ClienteService.ObtenerClientes(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, clientes)
}

func (h *Handler) obtenerClientesPorUsuario(ctx *gin.Context) {
	usuarioIdParam := ctx.Param("usuarioId")

	usuarioId, err := strconv.ParseInt(usuarioIdParam, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	clientes, err := h.services.ClienteService.ObtenerClientesPorUsuario(ctx.Request.Context(), usuarioId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, clientes)
}

func (h *Handler) obtenerClientesMovil(ctx *gin.Context) {
	usuarioId := ctx.GetInt64("usuarioId")
	fecha := ctx.Query("fecha")
	clientes, err := h.services.ClienteService.ObtenerClientesPorUsuarioMovil(ctx.Request.Context(), usuarioId, fecha)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, clientes)
}

func (h *Handler) crearCliente(ctx *gin.Context) {
	var clienteJSON model.CreateClienteModel

	clienteJSON.UsuarioCrea = ctx.GetInt64("usuarioId")
	clienteJSON.FechaCrea = time.Now()

	if err := ctx.BindJSON(&clienteJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	//usuarioId := ctx.GetInt64("usuarioId")
	nuevoCliente, err := h.services.ClienteService.CrearCliente(ctx.Request.Context(), clienteJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevoCliente)
}

func (h *Handler) actualizarCliente(ctx *gin.Context) {
	clienteIdParam := ctx.Param("clienteId")

	clienteId, err := strconv.ParseInt(clienteIdParam, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var clienteJSON model.UpdateClienteModel

	clienteJSON.UsuarioModifica = ctx.GetInt64("usuarioId")
	clienteJSON.FechaModifica = time.Now()

	if err := ctx.BindJSON(&clienteJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	actualizado, err := h.services.ClienteService.ActualizarCliente(ctx.Request.Context(), clienteId, clienteJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !actualizado {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo actualizar el cliente"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) validarCodigoClienteNuevo(ctx *gin.Context) {
	codigoCliente := ctx.Query("codigoCliente")
	cliente, err := h.services.ClienteService.ValidarCodigoClienteNuevo(ctx.Request.Context(), codigoCliente)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}

func (h *Handler) validarCodigoClienteModificar(ctx *gin.Context) {
	codigoCliente := ctx.Query("codigoCliente")
	id := ctx.Query("id")

	clienteId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	cliente, err := h.services.ClienteService.ValidarCodigoClienteModificar(codigoCliente, clienteId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}
