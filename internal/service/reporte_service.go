package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ahmetb/go-linq"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ReporteService interface {
	ObtenerImpulsadorasSubcidioTelefono(context.Context) (model.ImpulsadorasPayRollDataModel, error)
}

type reporteServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
}

func newReporteService(usuarioRepository repository.UsuarioRepository) *reporteServiceImpl {
	return &reporteServiceImpl{usuarioRepository: usuarioRepository}
}

func (c *reporteServiceImpl) ObtenerImpulsadorasSubcidioTelefono(ctx context.Context) (model.ImpulsadorasPayRollDataModel, error) {
	dataReporte := model.ImpulsadorasPayRollDataModel{}

	codigosImpulsadoras, err := c.usuarioRepository.ObtenerCodigoUsuariosImpulsadoras(ctx)
	if err != nil {
		return dataReporte, err
	}

	dataPayRoll, err := c.ObtenerImpulsadorasPayRoll(ctx)
	if err != nil {
		return dataReporte, err
	}

	for _, codigo := range codigosImpulsadoras {
		registroEncontrado := model.ImpulsadorasPayRollModel{}
		encontrado := linq.From(dataPayRoll).
			Where(func(c interface{}) bool {
				return c.(model.ImpulsadorasPayRollModel).Codigo == codigo.CodigoUsuario
			}).
			Any()

		if encontrado {
			registroEncontrado = linq.From(dataPayRoll).
				FirstWith(func(c interface{}) bool {
					return c.(model.ImpulsadorasPayRollModel).Codigo == codigo.CodigoUsuario
				}).(model.ImpulsadorasPayRollModel)

			if registroEncontrado.Codigo == "SIN CODIGO" {
				dataReporte.SinCodigoPayRoll = append(dataReporte.SinCodigoPayRoll, registroEncontrado)
			} else {

				registroEncontrado.Valor = "250.00"
				registroEncontrado.TipoContrato = codigo.TipoContrato

				dataReporte.Reporte = append(dataReporte.Reporte, registroEncontrado)
			}
		} else {
			estado, err := c.ObtenerEstadoImpulsadoraPayRoll(ctx, codigo.CodigoUsuario)
			if err != nil {
				return dataReporte, err
			}

			registroEncontrado.Valor = "0.00"
			registroEncontrado.Codigo = codigo.CodigoUsuario
			registroEncontrado.Nombre = codigo.Nombre
			registroEncontrado.Estado = estado
			dataReporte.NoEncontradosEnPayroll = append(dataReporte.NoEncontradosEnPayroll, registroEncontrado)
		}

	}

	return dataReporte, nil
}

func (c *reporteServiceImpl) ObtenerImpulsadorasPayRoll(ctx context.Context) ([]model.ImpulsadorasPayRollModel, error) {

	impulsadoras := []model.ImpulsadorasPayRollModel{}
	driver := os.Getenv("DB_DRIVER")

	connectionString := c.ConexionPayRoll(ctx)

	// Abre la conexión a SQL Server
	db, err := sql.Open(driver, connectionString)

	if err != nil {
		log.Fatalf("Error al abrir la conexión a SQL Server: %v", err)
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, `select COALESCE(COD_ALTERNO, 'SIN CODIGO') as codigo,
									CONCAT(PRIMER_NOMBRE , ' ' ,  SEGUNDO_NOMBRE, ' ' , APE_PATERNO, ' ' , APE_MATERNO ) as nombre,
									NUM_CUENTA_BANCO_PAGO as numeroCuenta,
									TIP_ESTADO
									from PLA_PERSONAL where COD_CATEGORIA = 66 and TIP_ESTADO = 'AC' `)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		var usuario model.ImpulsadorasPayRollModel

		err := rows.Scan(&usuario.Codigo, &usuario.Nombre, &usuario.NumeroCuenta, &usuario.Estado)
		if err != nil {
			println(err.Error())
			return impulsadoras, err
		}

		impulsadoras = append(impulsadoras, usuario)
	}
	return impulsadoras, nil

}

func (c *reporteServiceImpl) ObtenerEstadoImpulsadoraPayRoll(ctx context.Context, codigo string) (string, error) {

	driver := os.Getenv("DB_DRIVER")
	connectionString := c.ConexionPayRoll(ctx)

	// Abre la conexión a SQL Server
	db, err := sql.Open(driver, connectionString)

	if err != nil {
		log.Fatalf("Error al abrir la conexión a SQL Server: %v", err)
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, `select TIP_ESTADO from PLA_PERSONAL where COD_ALTERNO = $1 `, codigo)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()
	defer db.Close()

	estato := "No encontrado"

	for rows.Next() {

		err := rows.Scan(&estato)

		if err != nil {
			println(err.Error())
			return "Error al buscar impulsadora", err
		}

	}
	return estato, nil

}

func (c *reporteServiceImpl) ConexionPayRoll(ctx context.Context) string {

	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)

	return connectionString

}
