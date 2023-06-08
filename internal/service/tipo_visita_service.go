package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type TipoVisitaService interface {
	ObtenerTiposVisita(context.Context) ([]model.TipoVisitaModel, error)
	ObtenerTiposVisitaActiva(context.Context) ([]model.TipoVisitaModel, error)
	CrearTipoVisita(context.Context, model.CreateTipoVisitaModel) (model.TipoVisitaModel, error)
	ActualizarTipoVisita(context.Context, int64, model.UpdateTipoVisitaModel) (bool, error)
	ValidarTipoVisitaNuevo(string) (int64, error)
	ValidarTipoVisitaModificar(string, int64) (int64, error)
}

type tipoVisitaServiceImpl struct {
	tipoVisitaRepository repository.TipoVisitaRepository
}

func newTipoVisitaService(tipoVisitaRepository repository.TipoVisitaRepository) *tipoVisitaServiceImpl {
	return &tipoVisitaServiceImpl{tipoVisitaRepository: tipoVisitaRepository}
}

func (t *tipoVisitaServiceImpl) ObtenerTiposVisita(ctx context.Context) ([]model.TipoVisitaModel, error) {
	return t.tipoVisitaRepository.ObtenerTiposVisita(ctx)
}

func (t *tipoVisitaServiceImpl) ObtenerTiposVisitaActiva(ctx context.Context) ([]model.TipoVisitaModel, error) {
	return t.tipoVisitaRepository.ObtenerTiposVisitaActiva(ctx)
}

func (t *tipoVisitaServiceImpl) CrearTipoVisita(ctx context.Context, tipoVisita model.CreateTipoVisitaModel) (model.TipoVisitaModel, error) {
	var nuevoTipoVisita model.TipoVisitaModel

	idGenerado, err := t.tipoVisitaRepository.CrearTipoVisita(ctx, tipoVisita)

	if err != nil {
		return nuevoTipoVisita, err
	}

	nuevoTipoVisita.ID = idGenerado
	nuevoTipoVisita.Nombre = tipoVisita.Nombre
	nuevoTipoVisita.Color = tipoVisita.Color
	nuevoTipoVisita.Activo = true
	nuevoTipoVisita.RequiereMeta = *tipoVisita.RequiereMeta
	nuevoTipoVisita.RequiereMetaLinea = *tipoVisita.RequiereMetaLinea
	nuevoTipoVisita.RequiereMetaSubLinea = *tipoVisita.RequiereMetaSubLinea

	return nuevoTipoVisita, nil
}

func (t *tipoVisitaServiceImpl) ActualizarTipoVisita(ctx context.Context, tipoVisitaId int64, tipo model.UpdateTipoVisitaModel) (bool, error) {
	return t.tipoVisitaRepository.ActualizarTipoVisita(ctx, tipoVisitaId, tipo)
}

func (t *tipoVisitaServiceImpl) ValidarTipoVisitaNuevo(tipoVisita string) (int64, error) {
	return t.tipoVisitaRepository.ValidarTipoVisitaNuevo(tipoVisita)
}

func (t *tipoVisitaServiceImpl) ValidarTipoVisitaModificar(tipoVisita string, id int64) (int64, error) {
	return t.tipoVisitaRepository.ValidarTipoVisitaModificar(tipoVisita, id)
}
