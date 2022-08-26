package service

import "github.com/frangar97/mobilecheck-backend/internal/repository"

type Service struct {
	AuthService       AuthService
	ClienteService    ClienteService
	TipoVisitaService TipoVisitaService
	UsuarioService    UsuarioService
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		AuthService:       newAuthService(repositories.UsuarioRepository),
		ClienteService:    newClienteService(repositories.ClienteRepository),
		TipoVisitaService: newTipoVisitaService(repositories.TipoVisitaRepository),
		UsuarioService:    newUsuarioService(repositories.UsuarioRepository),
	}
}
