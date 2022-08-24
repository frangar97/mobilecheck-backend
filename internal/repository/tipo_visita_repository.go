package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TipoVisitaRepository interface {
	ObtenerTiposVisita(context.Context) ([]model.TipoVisitaModel, error)
	CrearTipoVisita(context.Context, model.CreateTipoVisitaModel) (int64, error)
}

type tipoVisitaRepositoryImpl struct {
	db *sql.DB
}

func newTipoVisitaRepository(db *sql.DB) *tipoVisitaRepositoryImpl {
	return &tipoVisitaRepositoryImpl{db: db}
}

func (t *tipoVisitaRepositoryImpl) ObtenerTiposVisita(ctx context.Context) ([]model.TipoVisitaModel, error) {
	tiposVisita := []model.TipoVisitaModel{}

	rows, err := t.db.QueryContext(ctx, "SELECT id,nombre,color,activo FROM TipoVisita")

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

	err := t.db.QueryRowContext(ctx, "INSERT INTO TipoVisita(nombre,color,activo) VALUES ($1,$2,$3) RETURNING id", tipoVisita.Nombre, tipoVisita.Color, true).Scan(&idGenerado)

	return idGenerado, err
}
