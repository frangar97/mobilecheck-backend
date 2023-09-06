package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TipoContratoRepository interface {
	ObtenerTipoContrato(context.Context) ([]model.TipoContratoModel, error)
}

type tipoContratoRepositoryImpl struct {
	db *sql.DB
}

func newTipoContratoRepository(db *sql.DB) *tipoContratoRepositoryImpl {
	return &tipoContratoRepositoryImpl{
		db: db,
	}
}

func (c *tipoContratoRepositoryImpl) ObtenerTipoContrato(ctx context.Context) ([]model.TipoContratoModel, error) {
	tiposContratos := []model.TipoContratoModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  id,
			tipo
	FROM tipoContrato 
	`)

	if err != nil {
		return tiposContratos, err
	}

	defer rows.Close()

	for rows.Next() {
		var tipo model.TipoContratoModel

		err := rows.Scan(&tipo.Id, &tipo.TipoContrato)

		if err != nil {
			return tiposContratos, err
		}

		tiposContratos = append(tiposContratos, tipo)
	}

	return tiposContratos, nil
}
