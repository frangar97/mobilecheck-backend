package repository

import "database/sql"

type Repository struct {
	UsuarioRepository    UsuarioRepository
	TipoVisitaRepository TipoVisitaRepository
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository:    newUsuarioRepository(db),
		TipoVisitaRepository: newTipoVisitaRepository(db),
	}
}
