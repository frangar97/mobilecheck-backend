package service

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type TareaService interface {
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (model.TareaModelMovil, error)
	CrearTareaWeb(context.Context, model.CreateTareaModelWeb) (model.TareaModelWeb, error)
	ObtenerTareasWeb(context.Context, string, string) ([]model.TareaModelWeb, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
	ObtenerCantidadTareasUsuarioPorFecha(context.Context, string, string) ([]model.CantidadTareaPorUsuario, error)
	CompletarTarea(*gin.Context, model.CompletarTareaModel, int64) error
	CrearTareaMasivaWeb(context.Context, model.CreateTareaMasivaModelWeb) error
}

type tareaServiceImpl struct {
	tareaRepository  repository.TareaRepository
	visitaRepository repository.VisitaRepository
}

func newTareaService(tareaRepository repository.TareaRepository, visitaRepository repository.VisitaRepository) *tareaServiceImpl {
	return &tareaServiceImpl{
		tareaRepository:  tareaRepository,
		visitaRepository: visitaRepository,
	}
}

func (t *tareaServiceImpl) CrearTareaMovil(ctx context.Context, tareaCreate model.CreateTareaModelMovil, usuarioId int64) (model.TareaModelMovil, error) {
	idGenerado, err := t.tareaRepository.CrearTareaMovil(ctx, tareaCreate, usuarioId)
	if err != nil {
		return model.TareaModelMovil{}, err
	}

	return t.tareaRepository.ObtenerTareaPorIdMovil(ctx, idGenerado)
}

func (t *tareaServiceImpl) CrearTareaWeb(ctx context.Context, tareaCreate model.CreateTareaModelWeb) (model.TareaModelWeb, error) {
	idGenerado, err := t.tareaRepository.CrearTareaWeb(ctx, tareaCreate)
	if err != nil {
		return model.TareaModelWeb{}, err
	}

	return t.tareaRepository.ObtenerTareaPorIdWeb(ctx, idGenerado)
}

func (t *tareaServiceImpl) CrearTareaMasivaWeb(ctx context.Context, tareaCreate model.CreateTareaMasivaModelWeb) error {

	for _, fecha := range tareaCreate.Fechas {
		tareaModel := model.CreateTareaModelWeb{
			ClienteId:       tareaCreate.ClienteId,
			UsuarioId:       tareaCreate.UsuarioId,
			Meta:            tareaCreate.Meta,
			Fecha:           fecha,
			ImagenRequerida: tareaCreate.ImagenRequerida,
			TipoVisitaId:    tareaCreate.TipoVisitaId,
		}

		_, err := t.tareaRepository.CrearTareaWeb(ctx, tareaModel)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tareaServiceImpl) ObtenerTareasWeb(ctx context.Context, fechaInicio string, fechaFinal string) ([]model.TareaModelWeb, error) {
	return t.tareaRepository.ObtenerTareasWeb(ctx, fechaInicio, fechaFinal)
}

func (t *tareaServiceImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	return t.tareaRepository.ObtenerTareasDelDia(ctx, fecha, usuarioId)
}

func (t *tareaServiceImpl) ObtenerCantidadTareasUsuarioPorFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadTareaPorUsuario, error) {
	return t.tareaRepository.ObtenerCantidadTareasUsuarioPorFecha(ctx, fechaInicio, fechaFin)
}

func (t *tareaServiceImpl) CompletarTarea(ctx *gin.Context, tarea model.CompletarTareaModel, usuarioId int64) error {
	var urlImagen string

	if *tarea.ImagenRequerida {
		formfile, _, err := ctx.Request.FormFile("imagen")

		if err != nil {
			return err
		}

		cld, _ := cloudinary.NewFromParams("dzmgbv4qn", "676166561161436", "H7JuKbIvzimY1qQXqKhIHX3i-nM")
		resp, err := cld.Upload.Upload(ctx, formfile, uploader.UploadParams{Folder: "samples"})

		if err != nil {
			return err
		}

		urlImagen = resp.SecureURL
	}

	visita := model.CreateVisitaModel{
		Comentario:   tarea.Comentario,
		Latitud:      tarea.Latitud,
		Longitud:     tarea.Longitud,
		Fecha:        tarea.Fecha,
		ClienteId:    tarea.ClienteId,
		TipoVisitaId: tarea.TipoVisitaId,
	}

	visitaId, err := t.visitaRepository.CrearVisita(ctx.Request.Context(), visita, urlImagen, usuarioId)

	if err != nil {
		return err
	}

	_, err = t.tareaRepository.CompletarTarea(ctx, tarea.TareaId, visitaId)

	if err != nil {
		return err
	}

	return nil
}
