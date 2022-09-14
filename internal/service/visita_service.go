package service

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type VisitaService interface {
	CrearVisita(*gin.Context, model.CreateVisitaModel, int64) (model.VisitaModel, error)
	ObtenerVisitasPorUsuario(context.Context, int64) ([]model.VisitaModel, error)
}

type visitaServiceImpl struct {
	visitaRepository repository.VisitaRepository
}

func newVisitaService(visitaRepository repository.VisitaRepository) *visitaServiceImpl {
	return &visitaServiceImpl{
		visitaRepository: visitaRepository,
	}
}

func (v *visitaServiceImpl) ObtenerVisitasPorUsuario(ctx context.Context, usuarioId int64) ([]model.VisitaModel, error) {
	return v.visitaRepository.ObtenerVisitasPorUsuario(ctx, usuarioId)
}

func (v *visitaServiceImpl) CrearVisita(ctx *gin.Context, visita model.CreateVisitaModel, usuarioId int64) (model.VisitaModel, error) {
	var visitaDB model.VisitaModel
	formfile, _, err := ctx.Request.FormFile("imagen")

	if err != nil {
		return visitaDB, err
	}

	cld, _ := cloudinary.NewFromParams("dzmgbv4qn", "676166561161436", "H7JuKbIvzimY1qQXqKhIHX3i-nM")
	resp, err := cld.Upload.Upload(ctx, formfile, uploader.UploadParams{Folder: "samples"})

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
