package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type TipoContratoService interface {
	ObtenerTipoContrato(context.Context) ([]model.TipoContratoModel, error)
}

type tipoContratoServiceImpl struct {
	tipoContratoRepository repository.TipoContratoRepository
}

func newTipoContratoService(tipoContratoRepository repository.TipoContratoRepository) *tipoContratoServiceImpl {
	return &tipoContratoServiceImpl{tipoContratoRepository: tipoContratoRepository}
}

func (c *tipoContratoServiceImpl) ObtenerTipoContrato(ctx context.Context) ([]model.TipoContratoModel, error) {
	return c.tipoContratoRepository.ObtenerTipoContrato(ctx)
}
