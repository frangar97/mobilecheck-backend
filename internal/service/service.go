package service

import "github.com/frangar97/mobilecheck-backend/internal/repository"

type Service struct {
	UsuarioService UsuarioService
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		UsuarioService: NewUsuarioService(repositories.UsuarioRepository),
	}
}
