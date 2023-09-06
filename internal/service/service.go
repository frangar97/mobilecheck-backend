package service

import "github.com/frangar97/mobilecheck-backend/internal/repository"

type Service struct {
	AuthService       AuthService
	ClienteService    ClienteService
	TipoVisitaService TipoVisitaService
	UsuarioService    UsuarioService
	VisitaService     VisitaService
	TareaService      TareaService
	PaisService       PaisService
	AccesoService     AccesoService
	ReporteService    ReporteService
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		AuthService:       newAuthService(repositories.UsuarioRepository, repositories.AccesoRepository),
		ClienteService:    newClienteService(repositories.ClienteRepository, repositories.UsuarioRepository),
		TipoVisitaService: newTipoVisitaService(repositories.TipoVisitaRepository),
		UsuarioService:    newUsuarioService(repositories.UsuarioRepository),
		VisitaService:     newVisitaService(repositories.VisitaRepository),
		TareaService:      newTareaService(repositories.TareaRepository, repositories.VisitaRepository, repositories.ClienteRepository, repositories.UsuarioRepository, repositories.TipoVisitaRepository),
		PaisService:       newPaisService(repositories.PaisRepository),
		AccesoService:     newAccesoService(repositories.AccesoRepository),
		ReporteService:    newReporteService(repositories.UsuarioRepository),
	}
}
