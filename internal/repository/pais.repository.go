package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type PaisRepository interface {
	ObtenerPaises(context.Context) ([]model.PaisModel, error)
}

type paisRepositoryImpl struct {
	db *sql.DB
}

func newPaisRepository(db *sql.DB) *paisRepositoryImpl {
	return &paisRepositoryImpl{
		db: db,
	}
}

func (c *paisRepositoryImpl) ObtenerPaises(ctx context.Context) ([]model.PaisModel, error) {
	paises := []model.PaisModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  C.id,
			C.nombre,
			C.Abreviatura,
			C.Activo
	FROM Pais C
	`)

	if err != nil {
		return paises, err
	}

	defer rows.Close()

	for rows.Next() {
		var pais model.PaisModel

		err := rows.Scan(&pais.ID, &pais.Nombre, &pais.Abreviatura, &pais.Activo)

		if err != nil {
			return paises, err
		}

		paises = append(paises, pais)
	}

	return paises, nil
}
