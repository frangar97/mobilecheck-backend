package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type SubsidioImpulsadorasRepository interface {
	ObtenerImpulsadorasSubcidioTelefono(context.Context) ([]model.ImpulsadorasPayRollModel, error)
	ObtenerEstadoImpulsadoras(context.Context, string) (string, error)
	EliminarImpulsadorasSubcidio(context.Context) (int64, error)
	CrearImpulsadorasSubcidio(context.Context, model.ImpulsadorasPayRollModel) (int64, error)
}

type subsidioImpulsadorasRepositoryImpl struct {
	postgresDB *sql.DB
}

func newSubsidioImpulsadorasRepository(postgresDB *sql.DB) *subsidioImpulsadorasRepositoryImpl {
	return &subsidioImpulsadorasRepositoryImpl{
		postgresDB: postgresDB,
	}
}

func (c *subsidioImpulsadorasRepositoryImpl) ObtenerImpulsadorasSubcidioTelefono(ctx context.Context) ([]model.ImpulsadorasPayRollModel, error) {
	impulsadoras := []model.ImpulsadorasPayRollModel{}

	rows, err := c.postgresDB.QueryContext(ctx, `select codigo, nombre, numerocuenta, estado, tipocontrato, tipoCuenta, banco, correo  from subsidioimpulsadoras where estado = 'AC' `)
	if err != nil {
		return impulsadoras, err
	}

	defer rows.Close()

	for rows.Next() {
		var usuario model.ImpulsadorasPayRollModel

		err := rows.Scan(&usuario.Codigo, &usuario.Nombre, &usuario.NumeroCuenta, &usuario.Estado, &usuario.TipoCuenta, &usuario.TipoCuenta, &usuario.Banco, &usuario.Correo)
		if err != nil {
			return impulsadoras, err
		}

		if usuario.TipoCuenta == "AS" {
			usuario.TipoCuenta = "CA"
		} else {
			usuario.TipoCuenta = "NO"
		}

		impulsadoras = append(impulsadoras, usuario)
	}
	return impulsadoras, nil

}

func (c *subsidioImpulsadorasRepositoryImpl) ObtenerEstadoImpulsadoras(ctx context.Context, codigo string) (string, error) {

	rows, err := c.postgresDB.QueryContext(ctx, `select estado from subsidioimpulsadoras where codigo = $1 `, codigo)
	if err != nil {
		return "Error", err
	}

	defer rows.Close()

	estato := "No encontrado"

	for rows.Next() {

		err := rows.Scan(&estato)

		if err != nil {
			return "Error al buscar impulsadora", err
		}

	}
	return estato, nil
}

func (t *subsidioImpulsadorasRepositoryImpl) EliminarImpulsadorasSubcidio(ctx context.Context) (int64, error) {
	res, err := t.postgresDB.ExecContext(ctx, `delete from subsidioimpulsadoras`)

	if err != nil {
		return 0, nil
	}
	count, err := res.RowsAffected()

	return count, err
}

func (t *subsidioImpulsadorasRepositoryImpl) CrearImpulsadorasSubcidio(ctx context.Context, impulsadora model.ImpulsadorasPayRollModel) (int64, error) {

	res, err := t.postgresDB.ExecContext(ctx, "INSERT INTO subsidioimpulsadoras(codigo, nombre, numerocuenta, estado, tipocontrato, tipoCuenta, banco, correo) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", impulsadora.Codigo, impulsadora.Nombre, impulsadora.NumeroCuenta, impulsadora.Estado, impulsadora.TipoContrato, impulsadora.TipoCuenta, impulsadora.Banco, impulsadora.Correo)

	if err != nil {
		return 0, nil
	}

	count, err := res.RowsAffected()

	return count, err
}
