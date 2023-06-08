package v1

import (
	"net/http"
	"strconv"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerUsuarios(ctx *gin.Context) {
	usuarios, err := h.services.UsuarioService.ObtenerUsuarios(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}

func (h *Handler) crearUsuario(ctx *gin.Context) {
	var usuarioJSON model.CreateUsuarioModel

	if err := ctx.BindJSON(&usuarioJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	nuevoUsuario, err := h.services.UsuarioService.CrearUsuario(ctx.Request.Context(), usuarioJSON)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nuevoUsuario)

}

func (h *Handler) actualizarUsuario(ctx *gin.Context) {
	usuarioIdParam := ctx.Param("usuarioId")

	usuarioId, err := strconv.ParseInt(usuarioIdParam, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var usuarioJSON model.UpdateUsuarioModel

	if err := ctx.BindJSON(&usuarioJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	actualizado, err := h.services.UsuarioService.ActualizarUsuario(ctx.Request.Context(), usuarioId, usuarioJSON)

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

func (h *Handler) validarUsuarioNuevo(ctx *gin.Context) {
	usuario := ctx.Query("usuarioCodigo")

	cliente, err := h.services.UsuarioService.ValidarUsuarioNuevo(usuario)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}

func (h *Handler) validarUsuarioModificar(ctx *gin.Context) {
	usuario := ctx.Query("usuarioCodigo")
	id := ctx.Query("id")

	tipoVisitaId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	existe, err := h.services.UsuarioService.ValidarUsuarioModificar(usuario, tipoVisitaId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, existe)
}

// ===============================Usuarios Asesores ===============================

func (h *Handler) obtenerAsesores(ctx *gin.Context) {
	usuarios, err := h.services.UsuarioService.ObtenerAsesores(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}
