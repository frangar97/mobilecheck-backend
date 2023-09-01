package service

import (
	"context"

	"github.com/ahmetb/go-linq"
	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
)

type AccesoService interface {
	ObtenerMenuUsuario(context.Context, int64, bool) ([]model.MenuUsuarioDataModel, error)
	ObtenerAccesosMenuUsuario(context.Context, int64, bool) (model.MenuAsignadoUsuario, error)
	ObtenerAccesosPantallaUsuario(context.Context, int64, int64) (model.PantallaAsignadaUsuario, error)
	AsignarMenu(context.Context, model.AsignarMenuUsuarioModel) (int64, error)
}

type accesoServiceImpl struct {
	accesoRepository repository.AccesoRepository
}

func newAccesoService(accesoRepository repository.AccesoRepository) *accesoServiceImpl {
	return &accesoServiceImpl{accesoRepository: accesoRepository}
}

func (c *accesoServiceImpl) ObtenerMenuUsuario(ctx context.Context, usuarioId int64, dispositivo bool) ([]model.MenuUsuarioDataModel, error) {
	return c.accesoRepository.ObtenerMenuUsuario(ctx, usuarioId, dispositivo)
}

func (c *accesoServiceImpl) ObtenerAccesosMenuUsuario(ctx context.Context, usuarioId int64, dispositivo bool) (model.MenuAsignadoUsuario, error) {
	Accesos := model.MenuAsignadoUsuario{}
	opcionMenu, err := c.accesoRepository.ObtenerMenu(ctx, dispositivo)
	if err != nil {
		return Accesos, err
	}

	permisosDelUsuario, err := c.accesoRepository.ObtenerPermisosMenuUsuario(ctx, usuarioId, dispositivo)
	if err != nil {
		return Accesos, err
	}

	for _, menu := range opcionMenu {
		registroEncontrado := model.MenuOpcionUsuarioModel{}
		asignado := linq.From(permisosDelUsuario).
			Where(func(c interface{}) bool {
				return c.(model.MenuOpcionUsuarioModel).IdMenuOpcion == menu.ID
			}).
			Any()

		if asignado {
			registroEncontrado = linq.From(permisosDelUsuario).
				FirstWith(func(c interface{}) bool {
					return c.(model.MenuOpcionUsuarioModel).IdMenuOpcion == menu.ID
				}).(model.MenuOpcionUsuarioModel)

			if registroEncontrado.Activo {
				Accesos.MenuAsignado = append(Accesos.MenuAsignado, menu)
			} else {
				Accesos.MenuNoAsignado = append(Accesos.MenuNoAsignado, menu)
			}

		} else {
			Accesos.MenuNoAsignado = append(Accesos.MenuNoAsignado, menu)
		}
	}

	return Accesos, nil
}

func (c *accesoServiceImpl) ObtenerAccesosPantallaUsuario(ctx context.Context, usuarioId int64, idOpcionMenu int64) (model.PantallaAsignadaUsuario, error) {
	Accesos := model.PantallaAsignadaUsuario{}
	pantallas, err := c.accesoRepository.ObtenerPantallas(ctx, idOpcionMenu)
	if err != nil {
		return Accesos, err
	}

	pantallasDelUsuario, err := c.accesoRepository.ObtenerPermisosPantallaUsuario(ctx, usuarioId, idOpcionMenu)
	if err != nil {
		return Accesos, err
	}

	for _, pantalla := range pantallas {
		registroEncontrado := model.PantallanUsuarioModel{}
		asignado := linq.From(pantallasDelUsuario).
			Where(func(c interface{}) bool {
				return c.(model.PantallanUsuarioModel).IdPantalla == pantalla.ID
			}).
			Any()

		if asignado {
			registroEncontrado = linq.From(pantallasDelUsuario).
				FirstWith(func(c interface{}) bool {
					return c.(model.PantallanUsuarioModel).IdPantalla == pantalla.ID
				}).(model.PantallanUsuarioModel)

			if registroEncontrado.Activo {
				Accesos.PantallaAsignada = append(Accesos.PantallaAsignada, pantalla)
			} else {
				Accesos.PantallaNoAsignada = append(Accesos.PantallaNoAsignada, pantalla)
			}

		} else {
			Accesos.PantallaNoAsignada = append(Accesos.PantallaNoAsignada, pantalla)
		}
	}

	return Accesos, nil
}

func (t *accesoServiceImpl) AsignarMenu(ctx context.Context, opcionMenu model.AsignarMenuUsuarioModel) (int64, error) {

	existe, err := t.accesoRepository.VaidarOpcionMenuAsignado(ctx, opcionMenu.Idusuario, opcionMenu.Idmenuopcion)

	if err != nil {
		println(err.Error())
		return 0, err
	}

	if existe.Id > 0 {

		idActualizado, err := t.accesoRepository.ActualizarMenuAsignado(ctx, existe.Id, !*existe.Activo, opcionMenu.UsuarioAccion, opcionMenu.FechaAccion)

		if err != nil {
			return 0, err
		}

		return idActualizado, err

	}

	idGenerado, err := t.accesoRepository.AsignarMenu(ctx, opcionMenu)

	if err != nil {
		return 0, err
	}

	return idGenerado, err
}
