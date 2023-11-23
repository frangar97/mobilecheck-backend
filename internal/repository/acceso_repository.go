package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type AccesoRepository interface {
	ObtenerPantallas(context.Context) ([]model.PantallaModel, error)
	ObtenerObjetos(context.Context) ([]model.ObjetoModel, error)
	ObtenerModulos(context.Context, bool, bool) ([]model.ModuloModel, error)
	ObtenerPantallasAsignadasUsuario(context.Context, int64, int64, bool, bool) ([]model.AccesoPantallaModel, error)
	ObtenerPantallasPorModulo(context.Context, int64, bool, bool) ([]model.PantallaModel, error)

	//----
	AsignarPantalla(context.Context, model.CreateUpdateAccesoPantallaModel) (int64, error)
	ValidarPantallaAsignada(context.Context, int64, int64) (model.AccesoPantallaModel, error)
	ActualizarPantallaAsignada(context.Context, int64, model.CreateUpdateAccesoPantallaModel) (int64, error)

	//----
	ObtenerPantallasAccesos(context.Context, int64, bool, bool) ([]model.PantallaAccesoModel, error)
}

type accesoRepositoryImpl struct {
	db *sql.DB
}

func newAccesoRepository(db *sql.DB) *accesoRepositoryImpl {
	return &accesoRepositoryImpl{
		db: db,
	}
}

func (c *accesoRepositoryImpl) ObtenerModulos(ctx context.Context, movil bool, web bool) ([]model.ModuloModel, error) {
	modulos := []model.ModuloModel{}

	var dispositivoQuery = "select m.id, m.modulo from modulos m inner join pantalla p on p.idmodulo = m.id where m.activo = true"

	if movil && !web {
		dispositivoQuery = dispositivoQuery + " and p.movil = true "
	} else if web && !movil {
		dispositivoQuery = dispositivoQuery + " and p.web = true "
	} else if web && movil {
		dispositivoQuery = dispositivoQuery + " and p.web = true and  p.movil = true "
	} else {
		return nil, nil
	}

	rows, err := c.db.QueryContext(ctx, dispositivoQuery+" group by m.id, m.modulo")

	if err != nil {
		return modulos, err
	}

	defer rows.Close()

	for rows.Next() {
		var modulo model.ModuloModel

		err := rows.Scan(&modulo.Id, &modulo.Modulo)

		if err != nil {
			return modulos, err
		}

		modulos = append(modulos, modulo)
	}

	return modulos, nil
}

func (c *accesoRepositoryImpl) ObtenerPantallas(ctx context.Context) ([]model.PantallaModel, error) {
	pantallas := []model.PantallaModel{}

	rows, err := c.db.QueryContext(ctx, `select id, pantalla, movil, web  from pantalla`)

	if err != nil {
		return pantallas, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallaModel

		err := rows.Scan(&pantalla.Id, &pantalla.Pantalla, &pantalla.Movil, &pantalla.Web)

		if err != nil {
			return pantallas, err
		}

		pantallas = append(pantallas, pantalla)
	}

	return pantallas, nil
}

func (c *accesoRepositoryImpl) ObtenerObjetos(ctx context.Context) ([]model.ObjetoModel, error) {
	objetos := []model.ObjetoModel{}

	rows, err := c.db.QueryContext(ctx, `select id, idpantalla, objeto  from objeto`)

	if err != nil {
		return objetos, err
	}

	defer rows.Close()

	for rows.Next() {
		var objeto model.ObjetoModel

		err := rows.Scan(&objeto.Id, &objeto.IdPantalla, &objeto.Objeto)

		if err != nil {
			return objetos, err
		}

		objetos = append(objetos, objeto)
	}

	return objetos, nil
}

func (c *accesoRepositoryImpl) ObtenerPantallasAsignadasUsuario(ctx context.Context, usuarioId int64, moduloId int64, movil bool, web bool) ([]model.AccesoPantallaModel, error) {
	accesoPantallas := []model.AccesoPantallaModel{}

	var dispositivo = ""

	if movil && !web {
		dispositivo = " and p.movil = true"
	} else if web && !movil {
		dispositivo = " and p.web = true"
	} else if web && movil {
		dispositivo = " and p.web = true and p.movil = true"
	} else {
		return nil, nil
	}

	rows, err := c.db.QueryContext(ctx, `select
										ap.idpantalla,
										p.pantalla,
										ap.activo,
										p.web,
										p.movil
									from
									accesopantalla ap
									inner join pantalla p on
	                                p.id = ap.idpantalla
									where ap.idusuario = $1 and p.idmodulo = $2 `+dispositivo, usuarioId, moduloId)
	if err != nil {
		return accesoPantallas, err
	}
	defer rows.Close()
	for rows.Next() {
		var accesoPantalla model.AccesoPantallaModel
		err := rows.Scan(&accesoPantalla.IdPantalla, &accesoPantalla.Pantalla, &accesoPantalla.Activo, &accesoPantalla.Web, &accesoPantalla.Movil)
		if err != nil {
			return accesoPantallas, err
		}
		accesoPantallas = append(accesoPantallas, accesoPantalla)
	}
	return accesoPantallas, nil
}

//----------------- ACCESO PANTALLA -----------------------------------------

func (t *accesoRepositoryImpl) AsignarPantalla(ctx context.Context, pantalla model.CreateUpdateAccesoPantallaModel) (int64, error) {
	var idGenerado int64
	err := t.db.QueryRowContext(ctx, "INSERT INTO accesopantalla(idpantalla, idusuario, activo, usuariocrea, fechacrea) VALUES($1,$2,$3,$4,$5) RETURNING id", pantalla.IdPantalla, pantalla.IdUsuario, pantalla.Activo, pantalla.UsuarioCrea, pantalla.FechaCrea).Scan(&idGenerado)
	return idGenerado, err
}

func (t *accesoRepositoryImpl) ValidarPantallaAsignada(ctx context.Context, usuarioId int64, idPantalla int64) (model.AccesoPantallaModel, error) {

	var pantalla model.AccesoPantallaModel
	err := t.db.QueryRowContext(ctx, "select id, activo from accesopantalla where idpantalla = $1 and idusuario = $2 LIMIT 1", idPantalla, usuarioId).Scan(&pantalla.IdPantalla, &pantalla.Activo)

	if err != nil {
		if err == sql.ErrNoRows {
			pantalla.Activo = true
			return pantalla, nil
		}
		return pantalla, err
	}

	return pantalla, err
}

func (t *accesoRepositoryImpl) ActualizarPantallaAsignada(ctx context.Context, id int64, pantalla model.CreateUpdateAccesoPantallaModel) (int64, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE accesopantalla
		SET idpantalla=$1, 
			idusuario=$2, 
			activo=$3, 
			usuariomodifica=$4, 
			fechamodifica=$5
		WHERE id=$6
	`, pantalla.IdPantalla, pantalla.IdUsuario, pantalla.Activo, pantalla.UsuarioModifica, pantalla.FechaModifica, id)
	if err != nil {
		println("error actualizar " + err.Error())
		return 0, nil
	}
	count, err := res.RowsAffected()
	if count > 0 {
		return id, nil
	}
	return 0, err
}

func (c *accesoRepositoryImpl) ObtenerPantallasPorModulo(ctx context.Context, idModulo int64, movil bool, web bool) ([]model.PantallaModel, error) {
	pantallas := []model.PantallaModel{}

	var dispositivo = ""

	if movil && !web {
		dispositivo = " and movil = true"
	} else if web && !movil {
		dispositivo = " and web = true"
	} else if web && movil {
		dispositivo = " and web = true and movil = true"
	} else {
		return nil, nil
	}

	rows, err := c.db.QueryContext(ctx, `select id,pantalla, movil, web from pantalla where idmodulo = $1 `+dispositivo, idModulo)

	if err != nil {
		return pantallas, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallaModel

		err := rows.Scan(&pantalla.Id, &pantalla.Pantalla, &pantalla.Movil, &pantalla.Web)

		if err != nil {
			return pantallas, err
		}

		pantallas = append(pantallas, pantalla)
	}

	return pantallas, nil
}

// -----------------Accesos--------------------------
func (c *accesoRepositoryImpl) ObtenerPantallasAccesos(ctx context.Context, idUsuario int64, movil bool, web bool) ([]model.PantallaAccesoModel, error) {
	pantallas := []model.PantallaAccesoModel{}

	var dispositivo = ""

	if movil && !web {
		dispositivo = " and p.movil = true"
	} else if web && !movil {
		dispositivo = " and p.web = true"
	} else {
		return nil, nil
	}

	rows, err := c.db.QueryContext(ctx, `select p.pantalla  from  accesopantalla ap
	inner join pantalla p on p.id = ap.idpantalla 
	where ap.idusuario = $1 and ap.activo = true `+dispositivo, idUsuario)

	if err != nil {
		return pantallas, err
	}

	defer rows.Close()

	for rows.Next() {
		var pantalla model.PantallaAccesoModel

		err := rows.Scan(&pantalla.Pantalla)

		if err != nil {
			return pantallas, err
		}

		pantallas = append(pantallas, pantalla)
	}

	return pantallas, nil
}
