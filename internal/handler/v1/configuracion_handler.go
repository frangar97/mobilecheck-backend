package v1

import (
	"net/http"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) obtenerConfiguracionSubsidioTelefonia(ctx *gin.Context) {
	configuracion, err := h.services.ConfiguracionService.ObtenerConfiguracionSubsidioTelefonia(ctx.Request.Context())

	if err != nil {
		print(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, configuracion)
}

func (h *Handler) actualizarParametro(ctx *gin.Context) {
	print("cabo")
	var parametroJSON model.ConfiguracionSubcidioUpdateModel
	print(parametroJSON.Id)
	print(parametroJSON.Parametro)

	if err := ctx.BindJSON(&parametroJSON); err != nil {
		print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "los datos enviados no son validos"})
		return
	}

	actualizado, err := h.services.ConfiguracionService.ActualizarParametro(ctx.Request.Context(), parametroJSON)

	if err != nil {
		print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !actualizado {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No se pudo actualizar el parametro"})
		return
	}

	ctx.Status(http.StatusOK)
}
