package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type UsuarioRepository interface {
	ObtenerUsuarios(context.Context) ([]model.UsuarioModel, error)
	CrearUsuario(context.Context, model.CreateUsuarioModel) (int64, error)
	ObtenerPorUsuario(context.Context, string) (model.UsuarioModel, error)
	ObtenerPorId(context.Context, int64) (model.UsuarioModel, error)
	ActualizarUsuario(context.Context, int64, model.UpdateUsuarioModel) (bool, error)
	ObtenerAsesores(context.Context) ([]model.UsuarioModel, error)
	ObtenerUsuarioPorId(context.Context, int64) (bool, error)
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
	usuarios := []model.UsuarioModel{}

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

func (u *usuarioRepositoryImpl) ObtenerPorUsuario(ctx context.Context, usuario string) (model.UsuarioModel, error) {
	var usuarioModel model.UsuarioModel

	err := u.db.QueryRowContext(ctx, "SELECT id,nombre,apellido,telefono,email,activo,usuario,password,web,movil FROM Usuario WHERE usuario = $1 LIMIT 1", usuario).Scan(&usuarioModel.ID, &usuarioModel.Nombre, &usuarioModel.Apellido, &usuarioModel.Telefono, &usuarioModel.Email, &usuarioModel.Activo, &usuarioModel.Usuario, &usuarioModel.Password, &usuarioModel.Web, &usuarioModel.Movil)

	return usuarioModel, err
}

func (u *usuarioRepositoryImpl) ObtenerPorId(ctx context.Context, usuarioId int64) (model.UsuarioModel, error) {
	var usuarioModel model.UsuarioModel

	err := u.db.QueryRowContext(ctx, "SELECT id,nombre,apellido,telefono,email,activo,usuario,password,web,movil FROM Usuario WHERE id = $1 LIMIT 1", usuarioId).Scan(&usuarioModel.ID, &usuarioModel.Nombre, &usuarioModel.Apellido, &usuarioModel.Telefono, &usuarioModel.Email, &usuarioModel.Activo, &usuarioModel.Usuario, &usuarioModel.Password, &usuarioModel.Web, &usuarioModel.Movil)

	return usuarioModel, err
}

func (u *usuarioRepositoryImpl) ActualizarUsuario(ctx context.Context, usuarioId int64, usuario model.UpdateUsuarioModel) (bool, error) {
	res, err := u.db.ExecContext(ctx, `
		UPDATE Usuario
		SET	   nombre = $1,
			   apellido = $2,
			   telefono = $3,
			   email = $4,
			   activo = $5,
			   usuario = $6,
			   web = $7,
			   movil = $8
		WHERE id = $9
	`, usuario.Nombre, usuario.Apellido, usuario.Telefono, usuario.Email, usuario.Activo, usuario.Usuario, usuario.Web, usuario.Movil, usuarioId)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}

// ===============================Usuarios Asesores ===============================

func (u *usuarioRepositoryImpl) ObtenerAsesores(ctx context.Context) ([]model.UsuarioModel, error) {
	usuarios := []model.UsuarioModel{}

	rows, err := u.db.QueryContext(ctx, "SELECT id,nombre,apellido,telefono,email,activo,usuario,web,movil FROM Usuario where movil = true")

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

func (t *usuarioRepositoryImpl) ObtenerUsuarioPorId(ctx context.Context, usuarioId int64) (bool, error) {

	rows, err := t.db.QueryContext(ctx, `select id from usuario where id = $1`, usuarioId)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	existe := false

	for rows.Next() {

		existe = true
	}

	return existe, nil
}
