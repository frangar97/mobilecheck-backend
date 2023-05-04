package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ClienteService interface {
	ObtenerClientes(context.Context) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuario(context.Context, int64) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuarioMovil(context.Context, int64, string) ([]model.ClienteModel, error)
	CrearCliente(context.Context, model.CreateClienteModel) (model.ClienteModel, error)
	ActualizarCliente(context.Context, int64, model.UpdateClienteModel) (bool, error)
}

type clienteServiceImpl struct {
	clienteRepository repository.ClienteRepository
	usuarioRepository repository.UsuarioRepository
}

func newClienteService(clienteRepository repository.ClienteRepository, usuarioRepository repository.UsuarioRepository) *clienteServiceImpl {
	return &clienteServiceImpl{
		clienteRepository: clienteRepository,
		usuarioRepository: usuarioRepository,
	}
}

func (c *clienteServiceImpl) ObtenerClientes(ctx context.Context) ([]model.ClienteModel, error) {
	return c.clienteRepository.ObtenerClientes(ctx)
}

func (c *clienteServiceImpl) ObtenerClientesPorUsuario(ctx context.Context, usuarioId int64) ([]model.ClienteModel, error) {
	return c.clienteRepository.ObtenerClientesPorUsuario(ctx, usuarioId)
}

func (c *clienteServiceImpl) ObtenerClientesPorUsuarioMovil(ctx context.Context, usuarioId int64, fecha string) ([]model.ClienteModel, error) {
	return c.clienteRepository.ObtenerClientesPorUsuarioMovil(ctx, usuarioId, fecha)
}

func (c *clienteServiceImpl) CrearCliente(ctx context.Context, clienteModel model.CreateClienteModel) (model.ClienteModel, error) {
	var cliente model.ClienteModel

	idGenerado, err := c.clienteRepository.CrearCliente(ctx, clienteModel)

	if err != nil {
		return cliente, err
	}

	cliente.ID = idGenerado
	cliente.Nombre = clienteModel.Nombre
	cliente.Activo = true
	cliente.Direccion = clienteModel.Direccion
	cliente.Telefono = clienteModel.Telefono
	cliente.Latitud = clienteModel.Latitud
	cliente.Longitud = clienteModel.Longitud
	cliente.Email = clienteModel.Email

	return cliente, nil
}

func (c *clienteServiceImpl) ActualizarCliente(ctx context.Context, clienteId int64, cliente model.UpdateClienteModel) (bool, error) {
	return c.clienteRepository.ActualizarCliente(ctx, clienteId, cliente)
}
