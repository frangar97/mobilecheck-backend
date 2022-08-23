package service

import "github.com/frangar97/mobilecheck-backend/internal/repository"

type Service struct {
	AuthService    AuthService
	UsuarioService UsuarioService
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		AuthService:    NewAuthService(repositories.UsuarioRepository),
		UsuarioService: NewUsuarioService(repositories.UsuarioRepository),
	}
}
