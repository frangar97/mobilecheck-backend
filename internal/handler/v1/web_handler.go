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
				cliente.POST("", h.crearCliente)
				cliente.PUT("/:clienteId", h.actualizarCliente)
				//cliente.GET("/:usuarioId", h.obtenerClientesPorUsuario)
				//====================Validaciones====================
				cliente.GET("validarCodigoClienteNuevo", h.validarCodigoClienteNuevo)
				cliente.GET("validarCodigoClienteModificar", h.validarCodigoClienteModificar)
			}

			usuario := authenticated.Group("/usuario")
			{
				usuario.GET("", h.obtenerUsuarios)
				usuario.POST("", h.crearUsuario)
				usuario.PUT("/:usuarioId", h.actualizarUsuario)
				usuario.GET("asesores", h.obtenerAsesores)
				usuario.PUT("updatePassword", h.updatePassword)
				//====================Validaciones====================
				usuario.GET("validarUsuarioNuevo", h.validarUsuarioNuevo)
				usuario.GET("validarUsuarioModificar", h.validarUsuarioModificar)
			}

			tipoVisita := authenticated.Group("/tipovisita")
			{
				tipoVisita.GET("", h.obtenerTiposVisita)
				tipoVisita.POST("", h.crearTipoVisita)
				tipoVisita.PUT("/:tipoVisitaId", h.actualizarTipoVisita)
				//====================Validaciones====================
				tipoVisita.GET("validarTipoVisitaNuevo", h.validarTipoVisitaNuevo)
				tipoVisita.GET("validarTipoVisitaModificar", h.validarTipoVisitaModificar)
			}

			visita := authenticated.Group("/visita")
			{
				visita.GET("", h.obtenerVisitasWebRangoFecha)
				visita.GET("cantidadusuario", h.obtenerVisitasWebCantidadPorUsuarioRangoFecha)
				visita.GET("cantidadtipo", h.obtenerVisitasWebCantidadPorTipoRangoFecha)
				visita.GET("visitatarea/:id", h.obtenerVisitaTareaWeb)
			}

			tarea := authenticated.Group("/tarea")
			{
				tarea.GET("cantidadusuario", h.obtenerTareasWebCantidadPorUsuarioRangoFecha)
				tarea.POST("", h.crearTareaWeb)
				tarea.GET("", h.obtenerTareasWeb)
				tarea.POST("tareasmasivas", h.crearTareaMasivaWeb)
				tarea.POST("tareasmasivasexcel", h.crearTareaMasivaExcelWeb)
				tarea.GET("verificarTarea", h.verificarTarea)
				tarea.GET("validarTareaExcel", h.validarTareaExcel)
				tarea.GET("obtenerTareasHoras", h.obtenerTareasHorasWeb)
				tarea.DELETE("eliminarTareas", h.eliminarTareas)
				tarea.GET("TareasImpulsadorDia", h.obtenerTareasDelDiaMovil)
				tarea.POST("completarTareaWeb", h.completarTarea)
			}

			pais := authenticated.Group("/pais")
			{
				pais.GET("", h.obtenerPaises)
			}

			acceso := authenticated.Group("/acceso")
			{
				acceso.GET("menuAsignado", h.obtenerMenuUsuario)
				acceso.GET("obtenerAccesosMenuUsuario", h.obtenerAccesosMenuUsuario)
				acceso.GET("obtenerAccesosPantallaUsuario", h.obtenerAccesosPantallaUsuario)
				acceso.POST("asignarMenuUsuario", h.asignarMenuUsuario)
			}

			reporte := authenticated.Group("/reporte")
			{
				reporte.GET("ObtenerImpulsadorasSubcidioTelefono", h.ObtenerImpulsadorasSubcidioTelefono)
			}

			cargoUsuario := authenticated.Group("/cargoUsuario")
			{
				cargoUsuario.GET("", h.obtenerCargoUsuario)
			}

			tipoContrato := authenticated.Group("/tipoContrato")
			{
				tipoContrato.GET("", h.obtenerTipoContrato)
			}

			importarExportarData := authenticated.Group("/importarExportarData")
			{
				importarExportarData.GET("actualizarSubsidioTelefoniaImpulsadoras", h.actualizarSubsidioTelefoniaImpulsadoras)
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

	token, permisos, err := h.services.AuthService.LoginWeb(ctx.Request.Context(), credenciales)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "permisos": permisos})
}
