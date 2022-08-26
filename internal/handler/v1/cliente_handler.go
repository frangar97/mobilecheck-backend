package v1

import (
	"net/http"
	"strconv"

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
	clientes, err := h.services.ClienteService.ObtenerClientesPorUsuario(ctx.Request.Context(), usuarioId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, clientes)
}

func (h *Handler) crearCliente(ctx *gin.Context) {
	var clienteJSON model.CreateClienteModel

	if err := ctx.BindJSON(&clienteJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	usuarioId := ctx.GetInt64("usuarioId")
	nuevoCliente, err := h.services.ClienteService.CrearCliente(ctx.Request.Context(), clienteJSON, usuarioId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevoCliente)
}
