package service

import (
	"context"

	"github.com/ahmetb/go-linq"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ReporteService interface {
	ObtenerImpulsadorasSubcidioTelefono(context.Context) (model.ImpulsadorasPayRollDataModel, error)
}

type reporteServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
	reporteRepository repository.SubsidioImpulsadorasRepository
}

func newReporteService(reporteRepository repository.SubsidioImpulsadorasRepository, usuarioRepository repository.UsuarioRepository) *reporteServiceImpl {
	return &reporteServiceImpl{reporteRepository: reporteRepository, usuarioRepository: usuarioRepository}
}

func (c *reporteServiceImpl) ObtenerImpulsadorasSubcidioTelefono(ctx context.Context) (model.ImpulsadorasPayRollDataModel, error) {
	dataReporte := model.ImpulsadorasPayRollDataModel{}

	codigosImpulsadoras, err := c.usuarioRepository.ObtenerCodigoUsuariosImpulsadoras(ctx)
	if err != nil {
		return dataReporte, err
	}

	dataPayRoll, err := c.reporteRepository.ObtenerImpulsadorasSubcidioTelefono(ctx)
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

			registroEncontrado.TipoContrato = codigo.TipoContrato

			if registroEncontrado.Banco == "900003" {
				dataReporte.ReporteFicohsa = append(dataReporte.ReporteFicohsa, registroEncontrado)
			} else {
				dataReporte.ReporteOtroBanco = append(dataReporte.ReporteFicohsa, registroEncontrado)

			}

		} else {
			estado, err := c.reporteRepository.ObtenerEstadoImpulsadoras(ctx, codigo.CodigoUsuario)
			if err != nil {
				return dataReporte, err
			}

			registroEncontrado.Codigo = codigo.CodigoUsuario
			registroEncontrado.Nombre = codigo.Nombre
			registroEncontrado.Estado = estado
			dataReporte.NoEncontradosEnPayroll = append(dataReporte.NoEncontradosEnPayroll, registroEncontrado)
		}

	}

	var impulsadorasFiltradas []model.ImpulsadorasPayRollModel
	linq.From(dataPayRoll).
		Where(func(c interface{}) bool {
			return c.(model.ImpulsadorasPayRollModel).Codigo == "SIN CODIGO"
		}).
		ToSlice(&impulsadorasFiltradas)

	dataReporte.SinCodigoPayRoll = impulsadorasFiltradas

	return dataReporte, nil
}
