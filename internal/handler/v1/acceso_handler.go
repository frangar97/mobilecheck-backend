package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerModulos(ctx *gin.Context) {

	movil, err := strconv.ParseBool(ctx.Query("movil"))
	if err != nil {
		return
	}

	web, err := strconv.ParseBool(ctx.Query("web"))
	if err != nil {
		return
	}

	accesos, err := h.services.AccesoService.ObtenerModulos(ctx.Request.Context(), movil, web)

	if err != nil {
		println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}

func (h *Handler) obtenerAccesosPantallasUsuario(ctx *gin.Context) {

	movil, err := strconv.ParseBool(ctx.Query("movil"))
	if err != nil {
		return
	}

	web, err := strconv.ParseBool(ctx.Query("web"))
	if err != nil {
		return
	}

	idUsuario, err := strconv.ParseInt(ctx.Query("idUsuario"), 0, 0)
	if err != nil {
		return
	}

	idModulo, err := strconv.ParseInt(ctx.Query("idModulo"), 0, 0)
	if err != nil {
		return
	}

	accesos, err := h.services.AccesoService.ObtenerAccesosPantallasUsuario(ctx.Request.Context(), idUsuario, idModulo, movil, web)

	if err != nil {
		println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}

func (h *Handler) asignarPantallaUsuario(ctx *gin.Context) {
	var pantallaJSON model.CreateUpdateAccesoPantallaModel

	usuarioAccion := ctx.GetInt64("usuarioId")

	pantallaJSON.FechaCrea = time.Now()
	pantallaJSON.FechaModifica = time.Now()
	pantallaJSON.UsuarioCrea = usuarioAccion
	pantallaJSON.UsuarioModifica = usuarioAccion

	if err := ctx.BindJSON(&pantallaJSON); err != nil {
		print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	accesos, err := h.services.AccesoService.AsignarPantalla(ctx.Request.Context(), pantallaJSON, usuarioAccion)
	if err != nil {

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, accesos)
}

///--------Movil

func (h *Handler) obtenerAccesosWebPorMovil(ctx *gin.Context) {

	usuario := ctx.GetInt64("usuarioId")

	accesos, err := h.services.AccesoService.ObtenerAccesosWebPorMovil(ctx.Request.Context(), usuario, false, true)
	println("----Accesos-----")
	for _, cc := range accesos {
		println(cc.Pantalla)
	}
	println("-------------------")
	if err != nil {
		println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}
