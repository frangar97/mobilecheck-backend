package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type TareaService interface {
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (model.TareaModelMovil, error)
	CrearTareaWeb(context.Context, model.CreateTareaModelWeb) (model.TareaModelWeb, error)
	ObtenerTareasWeb(context.Context, string, string) ([]model.TareaModelWeb, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
	ObtenerCantidadTareasUsuarioPorFecha(context.Context, string, string) ([]model.CantidadTareaPorUsuario, error)
}

type tareaServiceImpl struct {
	tareaRepository repository.TareaRepository
}

func newTareaService(tareaRepository repository.TareaRepository) *tareaServiceImpl {
	return &tareaServiceImpl{
		tareaRepository: tareaRepository,
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

func (t *tareaServiceImpl) ObtenerTareasWeb(ctx context.Context, fechaInicio string, fechaFinal string) ([]model.TareaModelWeb, error) {
	return t.tareaRepository.ObtenerTareasWeb(ctx, fechaInicio, fechaFinal)
}

func (t *tareaServiceImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	return t.tareaRepository.ObtenerTareasDelDia(ctx, fecha, usuarioId)
}

func (t *tareaServiceImpl) ObtenerCantidadTareasUsuarioPorFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadTareaPorUsuario, error) {
	return t.tareaRepository.ObtenerCantidadTareasUsuarioPorFecha(ctx, fechaInicio, fechaFin)
}
