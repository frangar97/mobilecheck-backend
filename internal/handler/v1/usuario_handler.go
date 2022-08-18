package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsuarioRoutes(api *gin.RouterGroup) {
	users := api.Group("/usuario")
	{
		users.GET("", h.obtenerUsuarios)
		users.POST("", h.crearUsuario)
	}
}

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
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	nuevoUsuario, err := h.services.UsuarioService.CrearUsuario(ctx.Request.Context(), usuarioJSON)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, nuevoUsuario)

}
