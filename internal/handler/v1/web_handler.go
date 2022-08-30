package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initWebRoutes(c *gin.RouterGroup) {
	web := c.Group("/web")
	{
		web.POST("/login", h.LoginWeb)
		web.POST("/register", h.crearUsuario)

		authenticated := web.Group("/", h.webIdentity)
		{
			cliente := authenticated.Group("/cliente")
			{
				cliente.GET("", h.obtenerClientes)
				cliente.GET("/:usuarioId", h.obtenerClientesPorUsuario)
			}

			usuario := authenticated.Group("/usuario")
			{
				usuario.GET("", h.obtenerUsuarios)
				usuario.POST("", h.crearUsuario)
			}

			tipoVisita := authenticated.Group("/tipovisita")
			{
				tipoVisita.GET("", h.obtenerTiposVisita)
				tipoVisita.POST("", h.crearTipoVisita)
				tipoVisita.PUT("/:tipoVisitaId", h.actualizarTipoVisita)
			}
		}

	}
}

func (h *Handler) LoginWeb(ctx *gin.Context) {
	var credenciales model.AuthCredencialModel

	if err := ctx.BindJSON(&credenciales); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	token, err := h.services.AuthService.LoginWeb(ctx.Request.Context(), credenciales)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
