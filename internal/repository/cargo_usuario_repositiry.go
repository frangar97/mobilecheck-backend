package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type CargoUsuarioRepository interface {
	ObtenerCargoUsuario(context.Context) ([]model.CargoUsuarioModel, error)
}

type cargoUsuarioRepositoryImpl struct {
	db *sql.DB
}

func newCargoUsuarioRepository(db *sql.DB) *cargoUsuarioRepositoryImpl {
	return &cargoUsuarioRepositoryImpl{
		db: db,
	}
}

func (c *cargoUsuarioRepositoryImpl) ObtenerCargoUsuario(ctx context.Context) ([]model.CargoUsuarioModel, error) {
	cargos := []model.CargoUsuarioModel{}

	rows, err := c.db.QueryContext(ctx, `
	SELECT  C.id,
			C.cargo
	FROM cargoUSuario C
	`)

	if err != nil {
		return cargos, err
	}

	defer rows.Close()

	for rows.Next() {
		var cargo model.CargoUsuarioModel

		err := rows.Scan(&cargo.Id, &cargo.Cargo)

		if err != nil {
			return cargos, err
		}

		cargos = append(cargos, cargo)
	}

	return cargos, nil
}
