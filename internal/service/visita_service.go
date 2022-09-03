package service

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type VisitaService interface {
	CrearVisita(*gin.Context, model.CreateVisitaModel, int64) (int64, error)
}

type visitaServiceImpl struct {
	visitaRepository repository.VisitaRepository
}

func newVisitaService(visitaRepository repository.VisitaRepository) *visitaServiceImpl {
	return &visitaServiceImpl{
		visitaRepository: visitaRepository,
	}
}

func (v *visitaServiceImpl) CrearVisita(ctx *gin.Context, visita model.CreateVisitaModel, usuarioId int64) (int64, error) {
	formfile, _, err := ctx.Request.FormFile("imagen")

	if err != nil {
		return 0, err
	}

	cld, _ := cloudinary.NewFromParams("dzmgbv4qn", "676166561161436", "H7JuKbIvzimY1qQXqKhIHX3i-nM")
	resp, err := cld.Upload.Upload(ctx, formfile, uploader.UploadParams{Folder: "samples"})

	if err != nil {
		return 0, err
	}

	idGenerado, err := v.visitaRepository.CrearVisita(ctx.Request.Context(), visita, resp.SecureURL, usuarioId)

	return idGenerado, err
}
