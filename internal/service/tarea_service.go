package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type TareaService interface {
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (model.TareaModelMovil, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
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

func (t *tareaServiceImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	return t.tareaRepository.ObtenerTareasDelDia(ctx, fecha, usuarioId)
}
