package service

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type VisitaService interface {
	CrearVisita(*gin.Context, model.CreateVisitaModel, int64) (model.VisitaModel, error)
	ObtenerVisitasPorRangoFecha(context.Context, string, string) ([]model.VisitaModel, error)
	ObtenerVisitasPorUsuarioDelDia(context.Context, string, int64) ([]model.VisitaModel, error)
	ObtenerCantidadVisitaPorUsuario(context.Context, string, string) ([]model.CantidadVisitaPorUsuario, error)
	ObtenerCantidadVisitaPorTipo(context.Context, string, string) ([]model.CantidadVisitaPorTipo, error)
	ObtenerVisitaTarea(context.Context, int64) ([]model.VisitaTareaModel, error)
}

type visitaServiceImpl struct {
	visitaRepository repository.VisitaRepository
}

func newVisitaService(visitaRepository repository.VisitaRepository) *visitaServiceImpl {
	return &visitaServiceImpl{
		visitaRepository: visitaRepository,
	}
}

func (v *visitaServiceImpl) ObtenerVisitasPorRangoFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.VisitaModel, error) {
	return v.visitaRepository.ObtenerVisitasPorRangoFecha(ctx, fechaInicio, fechaFin)
}

func (v *visitaServiceImpl) ObtenerCantidadVisitaPorUsuario(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadVisitaPorUsuario, error) {
	return v.visitaRepository.ObtenerCantidadVisitaPorUsuario(ctx, fechaInicio, fechaFin)
}

func (v *visitaServiceImpl) ObtenerCantidadVisitaPorTipo(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadVisitaPorTipo, error) {
	return v.visitaRepository.ObtenerCantidadVisitaPorTipo(ctx, fechaInicio, fechaFin)
}

func (v *visitaServiceImpl) ObtenerVisitasPorUsuarioDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.VisitaModel, error) {
	return v.visitaRepository.ObtenerVisitasPorUsuarioDelDia(ctx, fecha, usuarioId)
}

func (v *visitaServiceImpl) CrearVisita(ctx *gin.Context, visita model.CreateVisitaModel, usuarioId int64) (model.VisitaModel, error) {
	var visitaDB model.VisitaModel
	formfile, _, err := ctx.Request.FormFile("imagen")

	if err != nil {
		return visitaDB, err
	}

	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDNAME"), os.Getenv("APIKEY"), os.Getenv("APISECRET"))
	resp, err := cld.Upload.Upload(ctx, formfile, uploader.UploadParams{Folder: os.Getenv("CLOUDFOLDER")})

	if err != nil {
		return visitaDB, err
	}

	idGenerado, err := v.visitaRepository.CrearVisita(ctx.Request.Context(), visita, resp.SecureURL, usuarioId)

	if err != nil {
		return visitaDB, err
	}

	visitaDB, err = v.visitaRepository.ObtenerVisitaPorId(ctx, idGenerado)

	return visitaDB, err
}

func (v *visitaServiceImpl) ObtenerVisitaTarea(ctx context.Context, idTarea int64) ([]model.VisitaTareaModel, error) {
	return v.visitaRepository.ObtenerVisitaTarea(ctx, idTarea)
}
