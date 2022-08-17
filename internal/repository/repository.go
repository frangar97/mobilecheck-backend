package repository

import "database/sql"

type Repository struct {
	UsuarioRepository UsuarioRepository
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository: newUsuarioRepository(db),
	}
}
