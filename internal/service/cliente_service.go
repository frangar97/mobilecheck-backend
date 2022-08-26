package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ClienteService interface {
	ObtenerClientes(context.Context) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuario(context.Context, int64) ([]model.ClienteModel, error)
}

type clienteServiceImpl struct {
	clienteRepository repository.ClienteRepository
}

func newClienteService(clienteRepository repository.ClienteRepository) *clienteServiceImpl {
	return &clienteServiceImpl{
		clienteRepository: clienteRepository,
	}
}

func (c *clienteServiceImpl) ObtenerClientes(ctx context.Context) ([]model.ClienteModel, error) {
	return c.clienteRepository.ObtenerClientes(ctx)
}

func (c *clienteServiceImpl) ObtenerClientesPorUsuario(ctx context.Context, usuarioId int64) ([]model.ClienteModel, error) {
	return c.clienteRepository.ObtenerClientesPorUsuario(ctx, usuarioId)
}
