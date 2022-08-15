package repository

import "database/sql"

type Repository struct{}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{}
}
