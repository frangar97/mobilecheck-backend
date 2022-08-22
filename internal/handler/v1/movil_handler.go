package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initMovilRoutes(c *gin.RouterGroup) {
	web := c.Group("/web")
	{
		usuario := web.Group("/usuario")
		{
			usuario.GET("", h.obtenerUsuarios)
			usuario.POST("", h.crearUsuario)
		}
	}
}
