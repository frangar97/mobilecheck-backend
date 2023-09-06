package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type AccesoRepository interface {
	ObtenerMenuUsuario(context.Context, int64, bool) ([]model.MenuUsuarioDataModel, error)
	ObtenerPantallaUsuario(context.Context, int64, int64) ([]model.PantallaUsuarioDataModel, error)
	ObtenerPermisosMenuUsuario(context.Context, int64, bool) ([]model.MenuOpcionUsuarioModel, error)
	ObtenerMenu(context.Context, bool) ([]model.MenuUsuarioDataModel, error)
	ObtenerPermisosPantallaUsuario(context.Context, int64, int64) ([]model.PantallanUsuarioModel, error)
	ObtenerPantallas(context.Context, int64) ([]model.PantallaUsuarioDataModel, error)
	AsignarMenu(context.Context, model.AsignarMenuUsuarioModel) (int64, error)
	AsignarPantalla(context.Context, model.CreatePantallaUsuarioModel) (int64, error)
	VaidarOpcionMenuAsignado(context.Context, int64, int64) (model.AsignarMenuUsuarioModel, error)
	ActualizarMenuAsignado(context.Context, int64, bool, int64, time.Time) (int64, error)
}

type accesoRepositoryImpl struct {
	db *sql.DB
}

func newAccesoRepository(db *sql.DB) *accesoRepositoryImpl {
	return &accesoRepositoryImpl{
		db: db,
	}
}

func (c *accesoRepositoryImpl) ObtenerMenuUsuario(ctx context.Context, usuarioId int64, dispositivo bool) ([]model.MenuUsuarioDataModel, error) {
	menuUsuario := []model.MenuUsuarioDataModel{}

	rows, err := c.db.QueryContext(ctx, `
			select mo.id, mo.nombre
			from menuopcion mo
			inner join menuusuario mu on mo.id = mu.idmenuopcion 
			where  mu.idusuario = $1 and mo.dispositivo = $2 and mu.activo = true
	`, usuarioId, dispositivo)

	if err != nil {
		return menuUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var menu model.MenuUsuarioDataModel

		err := rows.Scan(&menu.ID, &menu.Opcion)

		if err != nil {
			return menuUsuario, err
		}

		menuUsuario = append(menuUsuario, menu)
	}

	return menuUsuario, nil
}

func (c *accesoRepositoryImpl) ObtenerPantallaUsuario(ctx context.Context, usuarioId int64, menu int64) ([]model.PantallaUsuarioDataModel, error) {
	pantallaUsuario := []model.PantallaUsuarioDataModel{}

	rows, err := c.db.QueryContext(ctx, `
	select P.id, P.nombre from pantallausuario PU
	inner join pantalla P on p.id = PU.idpantalla
	where PU.idusuario = $1 and P.idopcionmenu = $2 and PU.activo = true
	`, usuarioId, menu)

	if err != nil {
		return pantallaUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallaUsuarioDataModel

		err := rows.Scan(&pantalla.ID, &pantalla.Pantalla)

		if err != nil {
			return pantallaUsuario, err
		}

		pantallaUsuario = append(pantallaUsuario, pantalla)
	}

	return pantallaUsuario, nil
}

// //-----Asignacion y modificacion de permisos
// -----Menu
func (c *accesoRepositoryImpl) ObtenerPermisosMenuUsuario(ctx context.Context, usuarioId int64, dispositivo bool) ([]model.MenuOpcionUsuarioModel, error) {
	menuUsuario := []model.MenuOpcionUsuarioModel{}

	rows, err := c.db.QueryContext(ctx, `
	select mu.idmenuopcion , mu.activo  from menuopcion mo
			inner join menuusuario mu on mo.id = mu.idmenuopcion 
			where  mu.idusuario = $1 and mo.dispositivo = $2
	`, usuarioId, dispositivo)

	if err != nil {
		return menuUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var menu model.MenuOpcionUsuarioModel

		err := rows.Scan(&menu.IdMenuOpcion, &menu.Activo)

		if err != nil {
			return menuUsuario, err
		}

		menuUsuario = append(menuUsuario, menu)
	}

	return menuUsuario, nil
}

func (c *accesoRepositoryImpl) ObtenerMenu(ctx context.Context, dispositivo bool) ([]model.MenuUsuarioDataModel, error) {
	menuOpcion := []model.MenuUsuarioDataModel{}

	rows, err := c.db.QueryContext(ctx, `
	select id, nombre from menuopcion where dispositivo = $1
	`, dispositivo)

	if err != nil {
		return menuOpcion, err
	}

	defer rows.Close()

	for rows.Next() {
		var menu model.MenuUsuarioDataModel

		err := rows.Scan(&menu.ID, &menu.Opcion)

		if err != nil {
			return menuOpcion, err
		}

		menuOpcion = append(menuOpcion, menu)
	}

	return menuOpcion, nil
}

///----Menu Fin

// /------Pantalla
func (c *accesoRepositoryImpl) ObtenerPermisosPantallaUsuario(ctx context.Context, usuarioId int64, idOpcionMenu int64) ([]model.PantallanUsuarioModel, error) {
	pantallaUsuario := []model.PantallanUsuarioModel{}

	rows, err := c.db.QueryContext(ctx, `
	select PU.idpantalla, PU.activo from pantalla P
	inner join pantallausuario PU on P.id = PU.idpantalla 
	where PU.idusuario = $1 and P.idopcionmenu = $2 
	`, usuarioId, idOpcionMenu)

	if err != nil {
		return pantallaUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallanUsuarioModel

		err := rows.Scan(&pantalla.IdPantalla, &pantalla.Activo)

		if err != nil {
			return pantallaUsuario, err
		}

		pantallaUsuario = append(pantallaUsuario, pantalla)
	}

	return pantallaUsuario, nil
}

func (c *accesoRepositoryImpl) ObtenerPantallas(ctx context.Context, idOpcionMenu int64) ([]model.PantallaUsuarioDataModel, error) {
	pantallaOpcion := []model.PantallaUsuarioDataModel{}

	rows, err := c.db.QueryContext(ctx, `
	select id, nombre from pantalla where idopcionmenu = $1
	`, idOpcionMenu)

	if err != nil {
		return pantallaOpcion, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallaUsuarioDataModel

		err := rows.Scan(&pantalla.ID, &pantalla.Pantalla)

		if err != nil {
			return pantallaOpcion, err
		}

		pantallaOpcion = append(pantallaOpcion, pantalla)
	}

	return pantallaOpcion, nil
}

func (t *accesoRepositoryImpl) AsignarMenu(ctx context.Context, menu model.AsignarMenuUsuarioModel) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO menuusuario(idusuario, idmenuopcion, activo, usuariocrea, fechacrea) VALUES($1,$2,$3,$4,$5) RETURNING id", menu.Idusuario, menu.Idmenuopcion, menu.Activo, menu.UsuarioAccion, menu.FechaAccion).Scan(&idGenerado)

	return idGenerado, err
}

func (t *accesoRepositoryImpl) ActualizarMenuAsignado(ctx context.Context, id int64, activo bool, usuarioModifica int64, fechaModifica time.Time) (int64, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE menuusuario
		SET	  activo = $1,
		usuariomodifica = $2,
		fechamodifica = $3
		WHERE id = $4
	`, activo, usuarioModifica, fechaModifica, id)

	if err != nil {
		return 0, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return id, nil
	}

	return 0, err
}

func (t *accesoRepositoryImpl) VaidarOpcionMenuAsignado(ctx context.Context, usuarioId int64, idMenuOpcino int64) (model.AsignarMenuUsuarioModel, error) {

	var menuOpcionModel model.AsignarMenuUsuarioModel

	err := t.db.QueryRowContext(ctx, "select * from menuusuario where idusuario = $1 and idmenuopcion = $2 limit 1", usuarioId, idMenuOpcino).Scan(&menuOpcionModel.Id, &menuOpcionModel.Activo)

	return menuOpcionModel, err
}

func (t *accesoRepositoryImpl) AsignarPantalla(ctx context.Context, pantalla model.CreatePantallaUsuarioModel) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO pantallausuario(idpantalla, idusuario, activo, usuariomodifica, fechamodifica) VALUES($1,$2,$3,$4,$5) RETURNING id", pantalla.Idusuario, pantalla.Idpantalla, pantalla.Activo, pantalla.UsuarioCrea, pantalla.FechaCrea).Scan(&idGenerado)

	return idGenerado, err
}
