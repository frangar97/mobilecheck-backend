package repository

import "database/sql"

type Repository struct {
	ClienteRepository    ClienteRepository
	UsuarioRepository    UsuarioRepository
	TipoVisitaRepository TipoVisitaRepository
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository:    newUsuarioRepository(db),
		TipoVisitaRepository: newTipoVisitaRepository(db),
		ClienteRepository:    newClienteRepository(db),
	}
}
