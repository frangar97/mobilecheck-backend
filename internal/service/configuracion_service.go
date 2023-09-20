package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ConfiguracionService interface {
	ObtenerConfiguracionSubsidioTelefonia(context.Context) ([]model.ConfiguracionSubcidioModel, error)
	ActualizarParametro(context.Context, model.ConfiguracionSubcidioUpdateModel) (bool, error)
}

type configuracionServiceImpl struct {
	configuracionRepository repository.ConfiguracionRepository
}

func newConfiguracionService(configuracionRepository repository.ConfiguracionRepository) *configuracionServiceImpl {
	return &configuracionServiceImpl{configuracionRepository: configuracionRepository}
}

func (c *configuracionServiceImpl) ObtenerConfiguracionSubsidioTelefonia(ctx context.Context) ([]model.ConfiguracionSubcidioModel, error) {
	return c.configuracionRepository.ObtenerConfiguracionSubsidioTelefonia(ctx)
}

func (c *configuracionServiceImpl) ActualizarParametro(ctx context.Context, parametro model.ConfiguracionSubcidioUpdateModel) (bool, error) {
	return c.configuracionRepository.ActualizarParametro(ctx, parametro)
}
