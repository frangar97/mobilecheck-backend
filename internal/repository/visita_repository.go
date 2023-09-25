package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type VisitaRepository interface {
	CrearVisita(context.Context, model.CreateVisitaModel, string, int64) (int64, error)
	ObtenerVisitasPorRangoFecha(context.Context, string, string) ([]model.VisitaModel, error)
	ObtenerVisitasPorUsuarioDelDia(context.Context, string, int64) ([]model.VisitaModel, error)
	ObtenerVisitaPorId(context.Context, int64) (model.VisitaModel, error)
	ObtenerCantidadVisitaPorUsuario(context.Context, string, string) ([]model.CantidadVisitaPorUsuario, error)
	ObtenerCantidadVisitaPorTipo(context.Context, string, string) ([]model.CantidadVisitaPorTipo, error)
	ObtenerVisitaTarea(context.Context, int64) ([]model.VisitaTareaModel, error)
	ActualizarVisitaImagen(context.Context, int64, string) (bool, error)
}

type visitaRepositoryImpl struct {
	db *sql.DB
}

func newVisitaRepository(db *sql.DB) *visitaRepositoryImpl {
	return &visitaRepositoryImpl{db: db}
}

func (v *visitaRepositoryImpl) CrearVisita(ctx context.Context, visita model.CreateVisitaModel, imagenUrl string, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := v.db.QueryRowContext(ctx, "INSERT INTO Visita(comentario,latitud,longitud,fecha,imagen,usuarioId,clienteId,meta,metalinea,metasublinea,ip) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id", visita.Comentario, visita.Latitud, visita.Longitud, visita.Fecha, imagenUrl, usuarioId, visita.ClienteId, visita.Meta, visita.MetaLinea, visita.MetaSubLinea, visita.Ip).Scan(&idGenerado)

	return idGenerado, err
}

func (v *visitaRepositoryImpl) ObtenerVisitasPorRangoFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.VisitaModel, error) {
	visitasUsuario := []model.VisitaModel{}

	rows, err := v.db.QueryContext(ctx, `
		SELECT	V.id,
				V.comentario,
				V.latitud,
				V.longitud,
				V.imagen,
				V.fecha,
				C.nombre,
				TV.nombre,
				TV.color
		FROM	Visita V
		INNER JOIN Cliente C ON V.clienteId = C.id
		INNER JOIN TipoVisita TV ON V.tipoVisitaId = TV.id
		WHERE V.fecha BETWEEN $1 AND $2
		ORDER BY V.fecha DESC
	`, fechaInicio, fechaFin)

	if err != nil {
		return visitasUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var visita model.VisitaModel

		err := rows.Scan(&visita.ID, &visita.Comentario, &visita.Latitud, &visita.Longitud, &visita.Imagen, &visita.Fecha, &visita.Cliente /*&visita.TipoVisita,*/, &visita.Color)

		if err != nil {
			return visitasUsuario, err
		}

		visitasUsuario = append(visitasUsuario, visita)
	}

	return visitasUsuario, nil
}

func (v *visitaRepositoryImpl) ObtenerVisitasPorUsuarioDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.VisitaModel, error) {
	visitasUsuario := []model.VisitaModel{}

	rows, err := v.db.QueryContext(ctx, `
		SELECT	V.id,
				V.comentario,
				V.latitud,
				V.longitud,
				V.imagen,
				V.fecha,
				C.nombre,
				TV.nombre,
				TV.color
		FROM	Visita V
		INNER JOIN Cliente C ON V.clienteId = C.id
		INNER JOIN tarea T ON T.visitaid  = V.id  
		INNER JOIN TipoVisita TV ON TV.id = T.tipovisitaid 
		WHERE V.usuarioId = $1
		AND   DATE(V.fecha) = $2
		ORDER BY V.fecha DESC
	`, usuarioId, fecha)

	if err != nil {
		return visitasUsuario, err
	}

	defer rows.Close()

	for rows.Next() {
		var visita model.VisitaModel

		err := rows.Scan(&visita.ID, &visita.Comentario, &visita.Latitud, &visita.Longitud, &visita.Imagen, &visita.Fecha, &visita.Cliente, &visita.TipoVisita, &visita.Color)

		if err != nil {
			return visitasUsuario, err
		}

		visitasUsuario = append(visitasUsuario, visita)
	}

	return visitasUsuario, nil
}

func (v *visitaRepositoryImpl) ObtenerVisitaPorId(ctx context.Context, visitaId int64) (model.VisitaModel, error) {
	var visita model.VisitaModel

	err := v.db.QueryRowContext(ctx, `
		SELECT	V.id,
				V.comentario,
				V.latitud,
				V.longitud,
				V.imagen,
				V.fecha,
				C.nombre,
				TV.nombre,
				TV.color
		FROM	Visita V
		INNER JOIN Cliente C ON V.clienteId = C.id
		INNER JOIN TipoVisita TV ON V.tipoVisitaId = TV.id
		WHERE V.id = $1
		LIMIT 1
	`, visitaId).Scan(&visita.ID, &visita.Comentario, &visita.Latitud, &visita.Longitud, &visita.Imagen, &visita.Fecha, &visita.Cliente /*&visita.TipoVisita,*/, &visita.Color)

	return visita, err
}

func (v *visitaRepositoryImpl) ObtenerCantidadVisitaPorUsuario(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadVisitaPorUsuario, error) {
	visitas := []model.CantidadVisitaPorUsuario{}

	rows, err := v.db.QueryContext(ctx, `
		SELECT  CONCAT(U.nombre,' ',U.apellido) nombre,
		COUNT(*) cantidad
		FROM    VISITA V
			INNER JOIN Usuario U ON V.usuarioId = U.id
		WHERE   V.FECHA BETWEEN $1 AND $2
		GROUP BY U.nombre,U.apellido
		ORDER BY cantidad
		LIMIT 5
	`, fechaInicio, fechaFin)

	if err != nil {
		return visitas, err
	}

	for rows.Next() {
		var visita model.CantidadVisitaPorUsuario

		err := rows.Scan(&visita.Nombre, &visita.Cantidad)
		if err != nil {
			return visitas, err
		}

		visitas = append(visitas, visita)
	}

	return visitas, nil
}

func (v *visitaRepositoryImpl) ObtenerCantidadVisitaPorTipo(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadVisitaPorTipo, error) {
	visitas := []model.CantidadVisitaPorTipo{}

	rows, err := v.db.QueryContext(ctx, `
	SELECT  TV.nombre,
			TV.color,
			Count(*) cantidad
	FROM    tarea T
	LEFT JOIN VISITA V  on T.visitaid  = V.id 
	INNER JOIN TipoVisita TV ON T.tipovisitaid  = TV.id
	WHERE   t.FECHA BETWEEN $1 AND $2
	GROUP BY TV.nombre,TV.color
	`, fechaInicio, fechaFin)

	if err != nil {
		return visitas, err
	}

	for rows.Next() {
		var visita model.CantidadVisitaPorTipo

		err := rows.Scan(&visita.Nombre, &visita.Color, &visita.Cantidad)
		if err != nil {
			return visitas, err
		}

		visitas = append(visitas, visita)
	}

	return visitas, nil
}

func (v *visitaRepositoryImpl) ObtenerVisitaTarea(ctx context.Context, idTarea int64) ([]model.VisitaTareaModel, error) {
	visitaTarea := []model.VisitaTareaModel{}

	rows, err := v.db.QueryContext(ctx, `
	select 
		V.id,
		C.nombre  cliente,
		V.comentario,
		V.latitud,
		V.longitud,
		V.imagen,
		TV.nombre  tipoVisita,	
		V.fecha,
		V.meta metaVisita,
		T.meta metaTarea,
		TV.requieremeta,
		COALESCE(T.metaLinea,'') metalineaTarea,
		COALESCE(V.metaLinea,'') metalineaVisita,
		COALESCE(T.metaSublinea,'') metaSublineaTarea,
		COALESCE(V.metaSublinea,'') metaSublineaVisita,
		COALESCE(C.latitud,0) latitudCliente,
		COALESCE(c.longitud,0) longitudCliente,
		TV.requieremetalinea,
		TV.requieremetasublinea 
	from visita v 
	inner join tarea t on t.visitaid = v.id 
	inner join cliente c on c.id = t.clienteid 
	inner join tipovisita tv on tv.id = T.tipovisitaid 
	where t.id = $1
	`, idTarea)

	if err != nil {
		return visitaTarea, err
	}

	defer rows.Close()

	for rows.Next() {
		var visita model.VisitaTareaModel

		err := rows.Scan(&visita.ID, &visita.Cliente, &visita.Comentario, &visita.Latitud, &visita.Longitud, &visita.Imagen, &visita.TipoVisita, &visita.Fecha, &visita.MetaVisita, &visita.MetaTarea, &visita.RequiereMeta, &visita.MetaLineaTarea, &visita.MetaLineaVisita, &visita.MetaSubLineaTarea, &visita.MetaSubLineaVisita, &visita.LatitudCliente, &visita.LongitudCliente, &visita.RequiereMetaLinea, &visita.RequiereMetaSubLinea)

		if err != nil {
			return visitaTarea, err
		}

		visitaTarea = append(visitaTarea, visita)
	}

	return visitaTarea, nil
}

func (t *visitaRepositoryImpl) ActualizarVisitaImagen(ctx context.Context, visitaId int64, imagen string) (bool, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE visita
		SET	  imagen = $1
		WHERE id = $2
	`, imagen, visitaId)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}
