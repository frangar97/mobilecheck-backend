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
}

type tipoVisitaRepositoryImpl struct {
	db *sql.DB
}

func newTipoVisitaRepository(db *sql.DB) *tipoVisitaRepositoryImpl {
	return &tipoVisitaRepositoryImpl{db: db}
}

func (t *tipoVisitaRepositoryImpl) ObtenerTiposVisita(ctx context.Context) ([]model.TipoVisitaModel, error) {
	tiposVisita := []model.TipoVisitaModel{}

	rows, err := t.db.QueryContext(ctx, "SELECT id,nombre,color,activo,requieremeta FROM TipoVisita")

	if err != nil {
		return tiposVisita, err
	}

	defer rows.Close()

	for rows.Next() {
		var tipoVisita model.TipoVisitaModel

		err := rows.Scan(&tipoVisita.ID, &tipoVisita.Nombre, &tipoVisita.Color, &tipoVisita.Activo, &tipoVisita.RequiereMeta)

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

	err := t.db.QueryRowContext(ctx, "INSERT INTO TipoVisita(nombre,color,activo,requieremeta) VALUES ($1,$2,$3,$4) RETURNING id", tipoVisita.Nombre, tipoVisita.Color, true, tipoVisita.RequiereMeta).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tipoVisitaRepositoryImpl) ActualizarTipoVisita(ctx context.Context, tipoVisitaId int64, tipo model.UpdateTipoVisitaModel) (bool, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE TipoVisita
		SET	   nombre = $1,
			   color = $2,
			   activo = $3,
			   requiereMeta = $4
		WHERE id = $5
	`, tipo.Nombre, tipo.Color, tipo.Activo, tipo.RequiereMeta, tipoVisitaId)

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
