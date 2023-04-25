package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioService interface {
	ObtenerUsuarios(context.Context) ([]model.UsuarioModel, error)
	CrearUsuario(context.Context, model.CreateUsuarioModel) (model.UsuarioModel, error)
	ActualizarUsuario(context.Context, int64, model.UpdateUsuarioModel) (bool, error)
	ObtenerAsesores(context.Context) ([]model.UsuarioModel, error)
}

type usuarioServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
}

func newUsuarioService(usuarioRepository repository.UsuarioRepository) *usuarioServiceImpl {
	return &usuarioServiceImpl{
		usuarioRepository: usuarioRepository,
	}
}

func (u *usuarioServiceImpl) ObtenerUsuarios(ctx context.Context) ([]model.UsuarioModel, error) {
	return u.usuarioRepository.ObtenerUsuarios(ctx)
}

func (u *usuarioServiceImpl) CrearUsuario(ctx context.Context, usuario model.CreateUsuarioModel) (model.UsuarioModel, error) {
	var nuevoUsuario model.UsuarioModel

	usuarioBD, err := u.usuarioRepository.ObtenerPorUsuario(ctx, usuario.Usuario)

	if err != nil && err != sql.ErrNoRows {
		return nuevoUsuario, err
	}

	if (model.UsuarioModel{}) != usuarioBD {
		return nuevoUsuario, fmt.Errorf("el usuario %s ya esta en uso", usuarioBD.Usuario)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)

	if err != nil {
		return nuevoUsuario, err
	}

	usuario.Password = string(hashPassword)

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
	nuevoUsuario.Web = *usuario.Web
	nuevoUsuario.Movil = *usuario.Movil

	return nuevoUsuario, nil
}

func (u *usuarioServiceImpl) ActualizarUsuario(ctx context.Context, usuarioId int64, usuario model.UpdateUsuarioModel) (bool, error) {
	return u.usuarioRepository.ActualizarUsuario(ctx, usuarioId, usuario)
}

// ===============================Usuarios Asesores ===============================

func (u *usuarioServiceImpl) ObtenerAsesores(ctx context.Context) ([]model.UsuarioModel, error) {
	return u.usuarioRepository.ObtenerAsesores(ctx)
}
