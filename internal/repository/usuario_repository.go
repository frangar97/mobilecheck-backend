package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type UsuarioRepository interface {
	ObtenerUsuarios(context.Context) ([]model.UsuarioModel, error)
	CrearUsuario(context.Context, model.CreateUsuarioModel) (int64, error)
}

type usuarioRepositoryImpl struct {
	db *sql.DB
}

func newUsuarioRepository(db *sql.DB) *usuarioRepositoryImpl {
	return &usuarioRepositoryImpl{
		db: db,
	}
}

func (u *usuarioRepositoryImpl) ObtenerUsuarios(ctx context.Context) ([]model.UsuarioModel, error) {
	var usuarios []model.UsuarioModel

	rows, err := u.db.QueryContext(ctx, "SELECT id,nombre,apellido,telefono,email,activo,usuario,web,movil FROM Usuario")

	if err != nil {
		return usuarios, err
	}

	defer rows.Close()

	for rows.Next() {
		var usuario model.UsuarioModel

		err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellido, &usuario.Telefono, &usuario.Email, &usuario.Activo, &usuario.Usuario, &usuario.Web, &usuario.Movil)

		if err != nil {
			return usuarios, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u *usuarioRepositoryImpl) CrearUsuario(ctx context.Context, usuario model.CreateUsuarioModel) (int64, error) {
	var idGenerado int64

	err := u.db.QueryRowContext(ctx, "INSERT INTO Usuario(nombre,apellido,telefono,email,activo,usuario,password,web,movil) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", usuario.Nombre, usuario.Apellido, usuario.Telefono, usuario.Email, true, usuario.Usuario, usuario.Password, usuario.Web, usuario.Movil).Scan(&idGenerado)

	return idGenerado, err
}
