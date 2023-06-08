package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TipoVisitaRepository interface {
	ObtenerTiposVisita(context.Context) ([]model.TipoVisitaModel, error)
	ObtenerTiposVisitaActiva(context.Context) ([]model.TipoVisitaModel, error)
	CrearTipoVisita(context.Context, model.CreateTipoVisitaModel) (int64, error)
	ActualizarTipoVisita(context.Context, int64, model.UpdateTipoVisitaModel) (bool, error)
	ObtenerTipoVisitaPorId(context.Context, int64) (bool, error)
	ValidarTipoVisitaNuevo(string) (int64, error)
	ValidarTipoVisitaModificar(string, int64) (int64, error)
}

type tipoVisitaRepositoryImpl struct {
	db *sql.DB
}

func newTipoVisitaRepository(db *sql.DB) *tipoVisitaRepositoryImpl {
	return &tipoVisitaRepositoryImpl{db: db}
}

func (t *tipoVisitaRepositoryImpl) ObtenerTiposVisita(ctx context.Context) ([]model.TipoVisitaModel, error) {
	tiposVisita := []model.TipoVisitaModel{}

	rows, err := t.db.QueryContext(ctx, "SELECT id,nombre,color,activo,requieremeta,requieremetalinea,requieremetasublinea FROM TipoVisita")

	if err != nil {
		return tiposVisita, err
	}

	defer rows.Close()

	for rows.Next() {
		var tipoVisita model.TipoVisitaModel

		err := rows.Scan(&tipoVisita.ID, &tipoVisita.Nombre, &tipoVisita.Color, &tipoVisita.Activo, &tipoVisita.RequiereMeta, &tipoVisita.RequiereMetaLinea, &tipoVisita.RequiereMetaSubLinea)

		if err != nil {
			return tiposVisita, err
		}

		tiposVisita = append(tiposVisita, tipoVisita)
	}

	return tiposVisita, nil
}

func (t *tipoVisitaRepositoryImpl) ObtenerTiposVisitaActiva(ctx context.Context) ([]model.TipoVisitaModel, error) {
	tiposVisita := []model.TipoVisitaModel{}

	rows, err := t.db.QueryContext(ctx, "SELECT id,nombre,color,activo FROM TipoVisita WHERE activo = true")

	if err != nil {
		return tiposVisita, err
	}

	defer rows.Close()

	for rows.Next() {
		var tipoVisita model.TipoVisitaModel

		err := rows.Scan(&tipoVisita.ID, &tipoVisita.Nombre, &tipoVisita.Color, &tipoVisita.Activo)

		if err != nil {
			return tiposVisita, err
		}

		tiposVisita = append(tiposVisita, tipoVisita)
	}

	return tiposVisita, nil
}

func (t *tipoVisitaRepositoryImpl) CrearTipoVisita(ctx context.Context, tipoVisita model.CreateTipoVisitaModel) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO TipoVisita(nombre,color,activo,requieremeta,requieremetalinea,requieremetasublinea,usuariocrea,fechacrea) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id", tipoVisita.Nombre, tipoVisita.Color, true, tipoVisita.RequiereMeta, tipoVisita.RequiereMetaLinea, tipoVisita.RequiereMetaSubLinea, tipoVisita.UsuarioCrea, tipoVisita.FechaCrea).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tipoVisitaRepositoryImpl) ActualizarTipoVisita(ctx context.Context, tipoVisitaId int64, tipo model.UpdateTipoVisitaModel) (bool, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE TipoVisita
		SET	   nombre = $1,
			   color = $2,
			   activo = $3,
			   requiereMeta = $4,
			   requieremetalinea = $5,
			   requieremetasublinea = $6,
			   usuariomodifica = $7,
			   fechamodifica = $8
		WHERE id = $9
	`, tipo.Nombre, tipo.Color, tipo.Activo, tipo.RequiereMeta, tipo.RequiereMetaLinea, tipo.RequiereMetaSubLinea, tipo.UsuarioModifica, tipo.FechaModifica, tipoVisitaId)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}

func (t *tipoVisitaRepositoryImpl) ObtenerTipoVisitaPorId(ctx context.Context, tipoVisitaId int64) (bool, error) {

	rows, err := t.db.QueryContext(ctx, `select id from tipovisita where id = $1`, tipoVisitaId)
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

func (t *tipoVisitaRepositoryImpl) ValidarTipoVisitaNuevo(tipoVisita string) (int64, error) {
	var count int64
	err := t.db.QueryRow(`select count(nombre) from tipovisita where nombre = $1`, tipoVisita).Scan(&count)
	if err != nil {
		return 1, err
	}
	return count, nil
}

func (t *tipoVisitaRepositoryImpl) ValidarTipoVisitaModificar(tipoVisita string, id int64) (int64, error) {
	var count int64
	err := t.db.QueryRow(`select count(nombre) from tipovisita where nombre = $1 and id != $2`, tipoVisita, id).Scan(&count)
	if err != nil {
		return 1, err
	}
	return count, nil
}
