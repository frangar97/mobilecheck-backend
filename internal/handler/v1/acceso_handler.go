package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerMenuUsuario(ctx *gin.Context) {
	usuarioId, err := strconv.ParseInt(ctx.Query("usuario"), 0, 0)
	if err != nil {
		return
	}

	dispositivo, err := strconv.ParseBool(ctx.Query("webMovil"))
	if err != nil {
		return
	}
	accesos, err := h.services.AccesoService.ObtenerMenuUsuario(ctx.Request.Context(), usuarioId, dispositivo)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}

func (h *Handler) obtenerAccesosMenuUsuario(ctx *gin.Context) {

	usuarioId, err := strconv.ParseInt(ctx.Query("usuario"), 0, 0)
	if err != nil {
		return
	}

	dispositivo, err := strconv.ParseBool(ctx.Query("webMovil"))
	if err != nil {
		return
	}

	accesos, err := h.services.AccesoService.ObtenerAccesosMenuUsuario(ctx.Request.Context(), usuarioId, dispositivo)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}

func (h *Handler) obtenerAccesosPantallaUsuario(ctx *gin.Context) {

	usuarioId, err := strconv.ParseInt(ctx.Query("usuario"), 0, 0)
	if err != nil {
		return
	}

	opcionMenu, err := strconv.ParseInt(ctx.Query("opcionMenu"), 0, 0)
	if err != nil {
		return
	}

	accesos, err := h.services.AccesoService.ObtenerAccesosPantallaUsuario(ctx.Request.Context(), usuarioId, opcionMenu)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}

func (h *Handler) asignarMenuUsuario(ctx *gin.Context) {
	var OpcionMenuJSON model.AsignarMenuUsuarioModel

	Idusuario, err := strconv.ParseInt(ctx.Query("idusuario"), 0, 0)
	if err != nil {
		return
	}

	Idmenuopcion, err := strconv.ParseInt(ctx.Query("idmenuopcion"), 0, 0)
	if err != nil {
		return
	}

	OpcionMenuJSON.Idmenuopcion = Idmenuopcion
	OpcionMenuJSON.Idusuario = Idusuario
	OpcionMenuJSON.UsuarioAccion = ctx.GetInt64("usuarioId")
	OpcionMenuJSON.FechaAccion = time.Now()

	accesos, err := h.services.AccesoService.AsignarMenu(ctx.Request.Context(), OpcionMenuJSON)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, accesos)
}
