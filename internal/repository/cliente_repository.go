package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type ClienteRepository interface {
	ObtenerClientes(context.Context) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuario(context.Context, int64) ([]model.ClienteModel, error)
	CrearCliente(context.Context, model.CreateClienteModel, int64) (int64, error)
}

type clienteRepositoryImpl struct {
	db *sql.DB
}

func newClienteRepository(db *sql.DB) *clienteRepositoryImpl {
	return &clienteRepositoryImpl{
		db: db,
	}
}

func (c *clienteRepositoryImpl) ObtenerClientes(ctx context.Context) ([]model.ClienteModel, error) {
	clientes := []model.ClienteModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  C.id,
			C.nombre,
			C.telefono,
			C.email,
			C.direccion,
			C.latitud,
			C.longitud,
			C.activo,
			Concat(U.nombre,' ',U.apellido) Usuario
	FROM Cliente C
		  INNER JOIN usuario U ON C.usuarioId = U.id
	`)

	if err != nil {
		return clientes, err
	}

	defer rows.Close()

	for rows.Next() {
		var cliente model.ClienteModel

		err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Telefono, &cliente.Email, &cliente.Direccion, &cliente.Latitud, &cliente.Longitud, &cliente.Activo, &cliente.Usuario)

		if err != nil {
			return clientes, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (c *clienteRepositoryImpl) ObtenerClientesPorUsuario(ctx context.Context, usuarioId int64) ([]model.ClienteModel, error) {
	clientes := []model.ClienteModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  C.id,
			C.nombre,
			C.telefono,
			C.email,
			C.direccion,
			C.latitud,
			C.longitud,
			C.activo,
			Concat(U.nombre,' ',U.apellido) Usuario
	FROM Cliente C
		  INNER JOIN usuario U ON C.usuarioId = U.id
	WHERE C.usuarioId = $1
	`, usuarioId)

	if err != nil {
		return clientes, err
	}

	defer rows.Close()

	for rows.Next() {
		var cliente model.ClienteModel

		err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Telefono, &cliente.Email, &cliente.Direccion, &cliente.Latitud, &cliente.Longitud, &cliente.Activo, &cliente.Usuario)

		if err != nil {
			return clientes, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (c *clienteRepositoryImpl) CrearCliente(ctx context.Context, clienteModel model.CreateClienteModel, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := c.db.QueryRowContext(ctx, "INSERT INTO Cliente(nombre,telefono,email,direccion,latitud,longitud,activo,usuarioId) values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING ID", clienteModel.Nombre, clienteModel.Telefono, clienteModel.Email, clienteModel.Direccion, clienteModel.Latitud, clienteModel.Longitud, true, usuarioId).Scan(&idGenerado)

	return idGenerado, err
}
