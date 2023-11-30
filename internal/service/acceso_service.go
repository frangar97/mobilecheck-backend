package service

import (
	"context"

	"github.com/ahmetb/go-linq"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type AccesoService interface {
	ObtenerPantallas(context.Context) ([]model.PantallaModel, error)
	ObtenerObjetos(context.Context) ([]model.ObjetoModel, error)
	ObtenerModulos(context.Context, bool, bool) ([]model.ModuloModel, error)
	ObtenerAccesosPantallasUsuario(context.Context, int64, int64, bool, bool) (model.AccesoPantallaUsuarioModel, error)
	AsignarPantalla(context.Context, model.CreateUpdateAccesoPantallaModel, int64) (int64, error)
	ObtenerAccesosWebPorMovil(context.Context, int64, bool, bool) ([]model.PantallaAccesoModel, error)
	//ObtenerPantallasAccesos(context.Context, int64) ([]model.PantallaAccesoModel, error)
	ValidarAccesoPorId(context.Context, int64, int64) (bool, string, error)
}

type accesoServiceImpl struct {
	accesoRepository repository.AccesoRepository
}

func newAccesoService(accesoRepository repository.AccesoRepository) *accesoServiceImpl {
	return &accesoServiceImpl{accesoRepository: accesoRepository}
}

func (c *accesoServiceImpl) ObtenerModulos(ctx context.Context, movil bool, web bool) ([]model.ModuloModel, error) {
	return c.accesoRepository.ObtenerModulos(ctx, movil, web)
}

func (c *accesoServiceImpl) ObtenerPantallas(ctx context.Context) ([]model.PantallaModel, error) {
	return c.accesoRepository.ObtenerPantallas(ctx)
}

func (c *accesoServiceImpl) ObtenerObjetos(ctx context.Context) ([]model.ObjetoModel, error) {
	return c.accesoRepository.ObtenerObjetos(ctx)
}

func (c *accesoServiceImpl) ObtenerAccesosPantallasUsuario(ctx context.Context, usuarioId int64, moduloId int64, movil bool, web bool) (model.AccesoPantallaUsuarioModel, error) {
	Accesos := model.AccesoPantallaUsuarioModel{}
	pantallas, err := c.accesoRepository.ObtenerPantallasPorModulo(ctx, moduloId, movil, web)
	if err != nil {
		return Accesos, err
	}
	pantallasAsignadas, err := c.accesoRepository.ObtenerPantallasAsignadasUsuario(ctx, usuarioId, moduloId, movil, web)
	if err != nil {
		return Accesos, err
	}
	for _, pantalla := range pantallas {

		var pant model.AccesoPantallaModel
		pant.IdPantalla = pantalla.Id
		pant.Pantalla = pantalla.Pantalla
		pant.Movil = *pantalla.Movil
		pant.Web = *pantalla.Web
		pant.Activo = false

		registroEncontrado := model.AccesoPantallaModel{}
		asignado := linq.From(pantallasAsignadas).
			Where(func(c interface{}) bool {
				return c.(model.AccesoPantallaModel).IdPantalla == pant.IdPantalla
			}).
			Any()
		if asignado {
			registroEncontrado = linq.From(pantallasAsignadas).
				FirstWith(func(c interface{}) bool {
					return c.(model.AccesoPantallaModel).IdPantalla == pant.IdPantalla
				}).(model.AccesoPantallaModel)
			pant.Activo = registroEncontrado.Activo
			if registroEncontrado.Activo {
				Accesos.PantallasAsignadas = append(Accesos.PantallasAsignadas, pant)
			} else {
				Accesos.PantallasNoAsignadas = append(Accesos.PantallasNoAsignadas, pant)
			}
		} else {
			Accesos.PantallasNoAsignadas = append(Accesos.PantallasNoAsignadas, pant)
		}
	}
	return Accesos, nil
}

func (t *accesoServiceImpl) AsignarPantalla(ctx context.Context, pantalla model.CreateUpdateAccesoPantallaModel, idUsuario int64) (int64, error) {
	accesoExiste, err := t.accesoRepository.ValidarPantallaAsignada(ctx, pantalla.IdUsuario, pantalla.IdPantalla)
	if err != nil {
		return 0, err
	}
	println("----- ")
	println(accesoExiste.IdPantalla)
	println(idUsuario)
	if accesoExiste.IdPantalla > 0 {
		idActualizado, err := t.accesoRepository.ActualizarPantallaAsignada(ctx, accesoExiste.IdPantalla, pantalla)
		if err != nil {
			return 0, err
		}
		return idActualizado, err
	}

	idGenerado, err := t.accesoRepository.AsignarPantalla(ctx, pantalla)
	if err != nil {
		return 0, err
	}
	return idGenerado, err
}

// func (c *accesoServiceImpl) ObtenerPantallasAccesos(ctx context.Context, idUsuario int64) ([]model.PantallaAccesoModel, error) {
// 	return c.accesoRepository.ObtenerPantallasAccesos(ctx, idUsuario)
// }

func (c *accesoServiceImpl) ObtenerAccesosWebPorMovil(ctx context.Context, idUsuario int64, movil bool, web bool) ([]model.PantallaAccesoModel, error) {
	return c.accesoRepository.ObtenerPantallasAccesos(ctx, idUsuario, movil, web)
}

func (c *accesoServiceImpl) ValidarAccesoPorId(ctx context.Context, idPantalla int64, idUsuario int64) (bool, string, error) {
	return c.accesoRepository.ValidarAccesoPorId(ctx, idPantalla, idUsuario)
}
