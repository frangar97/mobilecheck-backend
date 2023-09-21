package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type ImportarExportarDataRepository interface {
	ObtenerImpulsadorasPayRoll(context.Context) ([]model.ImpulsadorasPayRollModel, error)
}

type importarExportarDataRepositoryImpl struct {
	postgresDB  *sql.DB
	sqlserverDB *sql.DB
}

func newImportarExportarDataRepository(postgresDB *sql.DB, sqlserverDB *sql.DB) *importarExportarDataRepositoryImpl {
	return &importarExportarDataRepositoryImpl{
		postgresDB:  postgresDB,
		sqlserverDB: sqlserverDB,
	}
}

func (c *importarExportarDataRepositoryImpl) ObtenerImpulsadorasPayRoll(ctx context.Context) ([]model.ImpulsadorasPayRollModel, error) {
	impulsadoras := []model.ImpulsadorasPayRollModel{}

	rows, err := c.sqlserverDB.QueryContext(ctx, `SELECT Codigo, Nombre, COALESCE(NumeroCuenta, '-'), Estado, COALESCE(TipoContrato, '-') as TipoContrato, Banco, TipoCuenta, Correo FROM IM_Reporte_Impulsadoras_Subsidio`)
	if err != nil {
		return impulsadoras, err
	}

	defer rows.Close()

	for rows.Next() {
		var usuario model.ImpulsadorasPayRollModel

		err := rows.Scan(&usuario.Codigo, &usuario.Nombre, &usuario.NumeroCuenta, &usuario.Estado, &usuario.TipoContrato, &usuario.Banco, &usuario.TipoCuenta, &usuario.Correo)
		if err != nil {
			return impulsadoras, err
		}

		impulsadoras = append(impulsadoras, usuario)
	}
	return impulsadoras, nil

}
