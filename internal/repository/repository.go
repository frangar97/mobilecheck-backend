package repository

import "database/sql"

type Repository struct {
	ClienteRepository      ClienteRepository
	UsuarioRepository      UsuarioRepository
	TipoVisitaRepository   TipoVisitaRepository
	VisitaRepository       VisitaRepository
	TareaRepository        TareaRepository
	PaisRepository         PaisRepository
	AccesoRepository       AccesoRepository
	CargoUsuarioRepository CargoUsuarioRepository
	TipoContratoRepository TipoContratoRepository
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository:      newUsuarioRepository(db),
		TipoVisitaRepository:   newTipoVisitaRepository(db),
		ClienteRepository:      newClienteRepository(db),
		VisitaRepository:       newVisitaRepository(db),
		TareaRepository:        newTareaRepository(db),
		PaisRepository:         newPaisRepository(db),
		AccesoRepository:       newAccesoRepository(db),
		CargoUsuarioRepository: newCargoUsuarioRepository(db),
		TipoContratoRepository: newTipoContratoRepository(db),
	}
}
