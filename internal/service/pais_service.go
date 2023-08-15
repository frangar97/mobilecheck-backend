package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type PaisService interface {
	ObtenerPaises(context.Context) ([]model.PaisModel, error)
}

type paisServiceImpl struct {
	paisRepository repository.PaisRepository
}

func newPaisService(paisRepository repository.PaisRepository) *paisServiceImpl {
	return &paisServiceImpl{paisRepository: paisRepository}
}

func (c *paisServiceImpl) ObtenerPaises(ctx context.Context) ([]model.PaisModel, error) {
	return c.paisRepository.ObtenerPaises(ctx)
}
