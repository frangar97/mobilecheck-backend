package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type ConfiguracionRepository interface {
	ObtenerConfiguracionSubsidioTelefonia(context.Context) ([]model.ConfiguracionSubcidioModel, error)
	ActualizarParametro(context.Context, model.ConfiguracionSubcidioUpdateModel) (bool, error)
}

type configuracionRepositoryImpl struct {
	db *sql.DB
}

func newConfiguracionRepository(db *sql.DB) *configuracionRepositoryImpl {
	return &configuracionRepositoryImpl{
		db: db,
	}
}

func (c *configuracionRepositoryImpl) ObtenerConfiguracionSubsidioTelefonia(ctx context.Context) ([]model.ConfiguracionSubcidioModel, error) {
	configuraciones := []model.ConfiguracionSubcidioModel{}

	rows, err := c.db.QueryContext(ctx, `select id, nombre, parametro, maxlength, minlength from configuracion where configuraciontipoid = 1`)

	if err != nil {
		return configuraciones, err
	}

	defer rows.Close()

	for rows.Next() {
		var configuracion model.ConfiguracionSubcidioModel

		err := rows.Scan(&configuracion.Id, &configuracion.Nombre, &configuracion.Parametro, &configuracion.Maxlength, &configuracion.Minlength)

		if err != nil {
			return configuraciones, err
		}

		configuraciones = append(configuraciones, configuracion)
	}

	return configuraciones, nil
}

func (t *configuracionRepositoryImpl) ActualizarParametro(ctx context.Context, parametro model.ConfiguracionSubcidioUpdateModel) (bool, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE configuracion
		SET	   parametro = $1
		WHERE id = $2
	`, parametro.Parametro, parametro.Id)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}
