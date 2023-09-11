package repository

import "database/sql"

type Repository struct {
	ClienteRepository              ClienteRepository
	UsuarioRepository              UsuarioRepository
	TipoVisitaRepository           TipoVisitaRepository
	VisitaRepository               VisitaRepository
	TareaRepository                TareaRepository
	PaisRepository                 PaisRepository
	AccesoRepository               AccesoRepository
	CargoUsuarioRepository         CargoUsuarioRepository
	TipoContratoRepository         TipoContratoRepository
	SubsidioImpulsadorasRepository SubsidioImpulsadorasRepository
	ImportarExportarDataRepository ImportarExportarDataRepository
}

func NewRepositories(postgresDB *sql.DB, sqlserverDB *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository:              newUsuarioRepository(postgresDB),
		TipoVisitaRepository:           newTipoVisitaRepository(postgresDB),
		ClienteRepository:              newClienteRepository(postgresDB),
		VisitaRepository:               newVisitaRepository(postgresDB),
		TareaRepository:                newTareaRepository(postgresDB),
		PaisRepository:                 newPaisRepository(postgresDB),
		AccesoRepository:               newAccesoRepository(postgresDB),
		CargoUsuarioRepository:         newCargoUsuarioRepository(postgresDB),
		TipoContratoRepository:         newTipoContratoRepository(postgresDB),
		SubsidioImpulsadorasRepository: newSubsidioImpulsadorasRepository(postgresDB),
		ImportarExportarDataRepository: newImportarExportarDataRepository(postgresDB, sqlserverDB),
	}
}
