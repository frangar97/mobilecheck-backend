package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type CargoUsuarioService interface {
	ObtenerCargoUsuario(context.Context) ([]model.CargoUsuarioModel, error)
}

type cargoUsuarioServiceImpl struct {
	cargoUsuarioRepository repository.CargoUsuarioRepository
}

func newCargoUsuarioService(cargoUsuarioRepository repository.CargoUsuarioRepository) *cargoUsuarioServiceImpl {
	return &cargoUsuarioServiceImpl{cargoUsuarioRepository: cargoUsuarioRepository}
}

func (c *cargoUsuarioServiceImpl) ObtenerCargoUsuario(ctx context.Context) ([]model.CargoUsuarioModel, error) {
	return c.cargoUsuarioRepository.ObtenerCargoUsuario(ctx)
}
