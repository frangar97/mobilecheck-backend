package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type ClienteRepository interface {
	ObtenerClientes(context.Context) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuario(context.Context, int64) ([]model.ClienteModel, error)
	ObtenerClientesPorUsuarioMovil(context.Context, int64, string) ([]model.ClienteModel, error)
	CrearCliente(context.Context, model.CreateClienteModel) (int64, error)
	ActualizarCliente(context.Context, int64, model.UpdateClienteModel) (bool, error)
	ObtenerClientePorId(context.Context, int64) (bool, error)
	ValidarCodigoClienteNuevo(context.Context, string) (bool, error)
	ValidarCodigoClienteModificar(string, int64) (int64, error)
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
			C.codigocliente,
			C.nombre,
			C.telefono,
			C.email,
			C.direccion,
			C.latitud,
			C.longitud,
			C.activo
	FROM Cliente C
	`)

	if err != nil {
		return clientes, err
	}

	defer rows.Close()

	for rows.Next() {
		var cliente model.ClienteModel

		err := rows.Scan(&cliente.ID, &cliente.CodigoCliente, &cliente.Nombre, &cliente.Telefono, &cliente.Email, &cliente.Direccion, &cliente.Latitud, &cliente.Longitud, &cliente.Activo)

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
			Concat(U.nombre,' ',U.apellido) Usuario,
			C.usuarioId
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

func (c *clienteRepositoryImpl) ObtenerClientesPorUsuarioMovil(ctx context.Context, usuarioId int64, fecha string) ([]model.ClienteModel, error) {
	clientes := []model.ClienteModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  C.id,
			C.nombre,
			C.telefono,
			C.email,
			C.direccion,
			C.latitud,
			C.longitud,
			C.activo			
	FROM cliente C
	INNER JOIN tarea T ON T.clienteid = C.id 	
	WHERE T.usuarioid = $1 AND Date(T.fecha) = $2 AND C.activo = true
	`, usuarioId, fecha)

	if err != nil {
		return clientes, err
	}

	defer rows.Close()

	for rows.Next() {
		var cliente model.ClienteModel

		err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Telefono, &cliente.Email, &cliente.Direccion, &cliente.Latitud, &cliente.Longitud, &cliente.Activo)

		if err != nil {
			return clientes, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (c *clienteRepositoryImpl) CrearCliente(ctx context.Context, clienteModel model.CreateClienteModel) (int64, error) {
	var idGenerado int64

	err := c.db.QueryRowContext(ctx, "INSERT INTO Cliente(codigocliente,nombre,telefono,email,direccion,activo,latitud,longitud) values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING ID", clienteModel.CodigoCliente, clienteModel.Nombre, clienteModel.Telefono, clienteModel.Email, clienteModel.Direccion, clienteModel.Activo, clienteModel.Latitud, clienteModel.Longitud).Scan(&idGenerado)

	return idGenerado, err
}

func (c *clienteRepositoryImpl) ActualizarCliente(ctx context.Context, clienteId int64, cliente model.UpdateClienteModel) (bool, error) {
	res, err := c.db.ExecContext(ctx, `
		UPDATE Cliente
		SET	   codigocliente = $1,
			   nombre = $2,
			   telefono = $3,
			   email = $4,
			   direccion = $5,
			   activo = $6,
			   latitud = $7,
			   longitud = $8
		WHERE id = $9
	`, cliente.CodigoCliente, cliente.Nombre, cliente.Telefono, cliente.Email, cliente.Direccion, cliente.Activo, cliente.Latitud, cliente.Longitud, clienteId)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}

func (t *clienteRepositoryImpl) ObtenerClientePorId(ctx context.Context, clienteId int64) (bool, error) {

	rows, err := t.db.QueryContext(ctx, `select id from cliente where id = $1`, clienteId)
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

func (t *clienteRepositoryImpl) ValidarCodigoClienteNuevo(ctx context.Context, codigoCliente string) (bool, error) {

	rows, err := t.db.QueryContext(ctx, `select codigocliente from cliente where codigocliente = $1`, codigoCliente)
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

func (t *clienteRepositoryImpl) ValidarCodigoClienteModificar(codigoCliente string, id int64) (int64, error) {
	var count int64
	err := t.db.QueryRow(`select count(codigocliente) from cliente where codigocliente = $1 and id != $2`, codigoCliente, id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
