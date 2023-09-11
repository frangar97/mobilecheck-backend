package service

import (
	"context"

	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type ImportarExportarDataService interface {
	ActualizarDataImpulsadoras(context.Context) (string, error)
}

type importarExportarDataServiceImpl struct {
	importarExportarDataRepository repository.ImportarExportarDataRepository
	subsidioImpulsadorasRepository repository.SubsidioImpulsadorasRepository
}

func newImportarExportarDataServiceImplService(importarExportarDataRepository repository.ImportarExportarDataRepository, subsidioImpulsadorasRepository repository.SubsidioImpulsadorasRepository) *importarExportarDataServiceImpl {
	return &importarExportarDataServiceImpl{importarExportarDataRepository: importarExportarDataRepository, subsidioImpulsadorasRepository: subsidioImpulsadorasRepository}
}

func (c *importarExportarDataServiceImpl) ActualizarDataImpulsadoras(ctx context.Context) (string, error) {
	_, err := c.subsidioImpulsadorasRepository.EliminarImpulsadorasSubcidio(ctx)

	if err != nil {
		return "Error al limpiar la data subsidio impulsadoras", nil
	}

	impulsadorasPayRoll, err := c.importarExportarDataRepository.ObtenerImpulsadorasPayRoll(ctx)

	if err != nil {
		return "Error al obtener data subsidio impulsadoras PAYROll", nil
	}

	for _, impulsadora := range impulsadorasPayRoll {
		_, err := c.subsidioImpulsadorasRepository.CrearImpulsadorasSubcidio(ctx, impulsadora)

		if err != nil {
			registroError := "Codigo" + impulsadora.Codigo + " Nombre" + impulsadora.Nombre
			return "Error en el registro" + registroError, nil
		}
	}
	return "Data actualizada", nil
}
