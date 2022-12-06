package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initMovilRoutes(c *gin.RouterGroup) {
	movil := c.Group("/movil")
	{
		movil.POST("/login", h.LoginMovil)
		movil.POST("/register", h.crearUsuario)

		authenticated := movil.Group("/", h.movilIdentity)
		{
			cliente := authenticated.Group("/cliente")
			{
				cliente.GET("", h.obtenerClientesMovil)
				cliente.POST("", h.crearCliente)
				cliente.PUT("/:clienteId", h.actualizarCliente)
			}

			tipoVisita := authenticated.Group("/tipovisita")
			{
				tipoVisita.GET("", h.obtenerTiposVisitaActiva)
			}

			visita := authenticated.Group("/visita")
			{
				visita.GET("", h.obtenerVisitasMovil)
				visita.POST("", h.CrearVisita)
			}

			tarea := authenticated.Group("/tarea")
			{
				tarea.GET("", h.obtenerTareasDelDiaMovil)
				tarea.POST("", h.crearTareaMovil)
				tarea.POST("/completar", h.completarTarea)
			}
		}
	}
}

func (h *Handler) LoginMovil(ctx *gin.Context) {
	var credenciales model.AuthCredencialModel

	if err := ctx.BindJSON(&credenciales); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	token, usuario, err := h.services.AuthService.LoginMovil(ctx.Request.Context(), credenciales)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "usuario": usuario})
}
