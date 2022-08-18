package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type UsuarioService interface {
	ObtenerUsuarios(context.Context) ([]model.UsuarioModel, error)
	CrearUsuario(context.Context, model.CreateUsuarioModel) (model.UsuarioModel, error)
}

type usuarioServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
}

func NewUsuarioService(usuarioRepository repository.UsuarioRepository) *usuarioServiceImpl {
	return &usuarioServiceImpl{
		usuarioRepository: usuarioRepository,
	}
}

func (u *usuarioServiceImpl) ObtenerUsuarios(ctx context.Context) ([]model.UsuarioModel, error) {
	return u.usuarioRepository.ObtenerUsuarios(ctx)
}

func (u *usuarioServiceImpl) CrearUsuario(ctx context.Context, usuario model.CreateUsuarioModel) (model.UsuarioModel, error) {
	var nuevoUsuario model.UsuarioModel

	idGenerado, err := u.usuarioRepository.CrearUsuario(ctx, usuario)

	if err != nil {
		return nuevoUsuario, err
	}

	nuevoUsuario.ID = idGenerado
	nuevoUsuario.Nombre = usuario.Usuario
	nuevoUsuario.Apellido = usuario.Apellido
	nuevoUsuario.Activo = true
	nuevoUsuario.Telefono = usuario.Telefono
	nuevoUsuario.Email = usuario.Email
	nuevoUsuario.Usuario = usuario.Usuario
	nuevoUsuario.Web = usuario.Web
	nuevoUsuario.Movil = usuario.Movil

	return nuevoUsuario, nil
}
