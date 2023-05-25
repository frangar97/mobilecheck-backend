package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TareaRepository interface {
	ObtenerTareaPorIdMovil(context.Context, int64) (model.TareaModelMovil, error)
	ObtenerTareaPorIdWeb(context.Context, int64) (model.TareaModelWeb, error)
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (int64, error)
	CrearTareaWeb(context.Context, model.CreateTareaModelWeb) (int64, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
	ObtenerTareasWeb(context.Context, string, string) ([]model.TareaModelWeb, error)
	ObtenerCantidadTareasUsuarioPorFecha(context.Context, string, string) ([]model.CantidadTareaPorUsuario, error)
	CompletarTarea(context.Context, int64, int64) (bool, error)
	VerificarTarea(context.Context, string, int64) (int, error)
	ObtenerTareasHorasWeb(context.Context, int, string) ([]model.TareaHorasModelReporteWeb, error)
}

type tareaRepositoryImpl struct {
	db *sql.DB
}

func newTareaRepository(db *sql.DB) *tareaRepositoryImpl {
	return &tareaRepositoryImpl{
		db: db,
	}
}

func (t *tareaRepositoryImpl) ObtenerTareaPorIdMovil(ctx context.Context, tareaId int64) (model.TareaModelMovil, error) {
	var tareaModel model.TareaModelMovil

	err := t.db.QueryRowContext(ctx, "SELECT T.id,T.Meta,T.fecha,T.completada,C.id,C.nombre FROM Tarea T INNER JOIN Cliente C ON T.clienteId = C.id WHERE T.id = $1 LIMIT 1", tareaId).Scan(&tareaModel.ID, &tareaModel.Meta, &tareaModel.Fecha, &tareaModel.Completada, &tareaModel.ClienteId, &tareaModel.Cliente)

	return tareaModel, err
}

func (t *tareaRepositoryImpl) ObtenerTareaPorIdWeb(ctx context.Context, tareaId int64) (model.TareaModelWeb, error) {
	var tareaModel model.TareaModelWeb

	err := t.db.QueryRowContext(ctx, "SELECT T.id,T.fecha,T.completada,C.nombre FROM Tarea T INNER JOIN CLIENTE C ON T.clienteId = C.id WHERE T.id = $1 LIMIT 1", tareaId).Scan(&tareaModel.ID, &tareaModel.Fecha, &tareaModel.Completada, &tareaModel.Cliente)

	return tareaModel, err
}

func (t *tareaRepositoryImpl) CrearTareaMovil(ctx context.Context, tarea model.CreateTareaModelMovil, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO Tarea(descripcion,fecha,clienteId,usuarioId,completada) VALUES($1,$2,$3,$4,false) RETURNING id", tarea.Descripcion, tarea.Fecha, tarea.ClienteId, usuarioId).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tareaRepositoryImpl) CrearTareaWeb(ctx context.Context, tarea model.CreateTareaModelWeb) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO Tarea(clienteId,usuarioId,tipovisitaid,meta,fecha,imagenRequerida,completada,metalinea,metasublinea) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", tarea.ClienteId, tarea.UsuarioId, tarea.TipoVisitaId, tarea.Meta, tarea.Fecha, tarea.ImagenRequerida, false, tarea.MetaLinea, tarea.MetaSubLinea).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tareaRepositoryImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	rows, err := t.db.QueryContext(ctx, `SELECT T.id,
				T.fecha,
				T.completada,
				C.id as clienteId,
				C.nombre as cliente,
				T.imagenrequerida,
				TV.nombre as tipoVisita,
				T.meta,
				TV.requieremeta  
			FROM Tarea T 
			INNER JOIN Cliente C ON T.clienteId = C.id
			inner join tipovisita TV on TV.id = T.tipovisitaid 
			WHERE DATE(T.fecha) = $1  AND T.usuarioId = $2 ORDER BY T.fecha`, fecha, usuarioId)
	if err != nil {
		return []model.TareaModelMovil{}, err
	}

	defer rows.Close()

	tareas := []model.TareaModelMovil{}

	for rows.Next() {
		var tarea model.TareaModelMovil

		err := rows.Scan(&tarea.ID, &tarea.Fecha, &tarea.Completada, &tarea.ClienteId, &tarea.Cliente, &tarea.ImagenRequerida, &tarea.TipoVisita, &tarea.Meta, &tarea.Requieremeta)
		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	return tareas, nil
}

func (t *tareaRepositoryImpl) ObtenerTareasWeb(ctx context.Context, fechaInicio string, fechaFinal string) ([]model.TareaModelWeb, error) {
	rows, err := t.db.QueryContext(ctx, `SELECT T.id,
				COALESCE(V.fecha, '0001-01-01 01:00:00+00') fecha,
				T.completada,
				C.nombre  cliente,
				T.imagenRequerida, 
				CONCAT(U.nombre,' ', U.apellido) asesor,	
				COALESCE(v.latitud,0) latitud,
				COALESCE(v.longitud,0) longitud,
				COALESCE(v.imagen,'') imagen,
				TV.nombre  tipoVisita,
				COALESCE(V.comentario,'') comentario,
				TV.requiereMeta,
				COALESCE(T.meta,'') metaAsignada,
				COALESCE(V.meta,'') metaCumplida,
				U.usuario codigoUsuario,
				U.id usuarioId,
				COALESCE(T.metalinea,'') metaLineaAsignada,
				COALESCE(T.metasublinea,'') metaSubLineaAsignada,
				COALESCE(V.metalinea,'') metaLineaCumplida,
				COALESCE(V.metasublinea,'') metaSubLineaCumpida,
				C.codigoCliente,
				COALESCE(C.latitud,0) latitudCliente,
				COALESCE(C.longitud,0) longitudCliente
			FROM Tarea T 
			INNER JOIN CLIENTE C ON T.clienteId = C.id 
			INNER JOIN USUARIO U ON T.usuarioid  = U.id
			LEFT  JOIN visita v on V.id = T.visitaid
			INNER JOIN tipovisita TV ON TV.id = T.tipovisitaid  
		WHERE T.fecha BETWEEN $1 AND $2 ORDER BY T.fecha`, fechaInicio, fechaFinal)
	if err != nil {
		return []model.TareaModelWeb{}, err
	}

	defer rows.Close()

	tareas := []model.TareaModelWeb{}

	for rows.Next() {
		var tarea model.TareaModelWeb

		err := rows.Scan(&tarea.ID, &tarea.Fecha, &tarea.Completada, &tarea.Cliente, &tarea.ImagenRequerida, &tarea.Asesor, &tarea.Latitud, &tarea.Longitud, &tarea.Imagen, &tarea.TipoVisita, &tarea.Comentario, &tarea.Requieremeta, &tarea.MetaAsignada, &tarea.MetaCumplida, &tarea.CodigoUsuario, &tarea.UsuarioId, &tarea.MetaLineaAsignada, &tarea.MetaSubLineaAsignada, &tarea.MetaLineaCumplida, &tarea.MetaSubLineaCumplida, &tarea.CodigoCliente, &tarea.LatitudCliente, &tarea.LongitudCliente)
		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	return tareas, nil
}

func (t *tareaRepositoryImpl) ObtenerCantidadTareasUsuarioPorFecha(ctx context.Context, fechaInicio string, fechaFin string) ([]model.CantidadTareaPorUsuario, error) {
	tareas := []model.CantidadTareaPorUsuario{}

	rows, err := t.db.QueryContext(ctx, `
	SELECT  CONCAT(U.nombre,' ',U.apellido) nombre,
        COUNT(*) filter (where T.completada = true) completadas,
        COUNT(*) filter (where T.completada = false) pendientes,
        COUNT(*) total
	FROM    Tarea T
			INNER JOIN Usuario U ON T.usuarioId = U.id
	WHERE   T.FECHA BETWEEN $1 AND $2
	GROUP BY U.nombre,U.apellido
	`, fechaInicio, fechaFin)

	if err != nil {
		return tareas, err
	}

	defer rows.Close()

	for rows.Next() {
		var tarea model.CantidadTareaPorUsuario

		err := rows.Scan(&tarea.Nombre, &tarea.Completadas, &tarea.Pendientes, &tarea.Total)
		if err != nil {
			return tareas, err
		}

		tareas = append(tareas, tarea)
	}

	return tareas, nil
}

func (t *tareaRepositoryImpl) CompletarTarea(ctx context.Context, tareaId int64, visitaId int64) (bool, error) {
	res, err := t.db.ExecContext(ctx, `
		UPDATE Tarea
		SET	  visitaId = $1,
			  completada = $2
		WHERE id = $3
	`, visitaId, true, tareaId)

	if err != nil {
		return false, nil
	}

	count, err := res.RowsAffected()

	if count > 0 {
		return true, nil
	}

	return false, err
}

func (t *tareaRepositoryImpl) VerificarTarea(ctx context.Context, fecha string, usuarioId int64) (int, error) {

	rows, err := t.db.QueryContext(ctx, `select * from tarea where fecha = $1 and usuarioid = $2`, fecha, usuarioId)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	tarea := 0

	for rows.Next() {

		tarea = 1
	}

	return tarea, nil
}

func (t *tareaRepositoryImpl) ObtenerTareasHorasWeb(ctx context.Context, usuarioId int, fecha string) ([]model.TareaHorasModelReporteWeb, error) {
	rows, err := t.db.QueryContext(ctx, `SELECT    a.codigousuario,
									a.asesor,
									a.codigoCliente,
									a.cliente,
									a.fechaTarea,
									a.fechaVisita as fechaEntrada,	
									b.fechaVisita as fechaSalida,	
									a.comentario as comentarioEntrada,
									b.comentario as comentarioSalida,
									CONCAT('https://maps.google.com/?q=',a.latitud,',',a.longitud)as ubicacionEntrada,
									CONCAT('https://maps.google.com/?q=',b.latitud,',',b.longitud)as ubicacionSalida,
									a.imagen as imagenEntrada,
									b.imagen as imaenSalida,
									a.latitud as latitudEntrada,
									a.longitud as longitudEntrada,
									a.latitud as latitudSalida,
									a.longitud as longitudSalida,
									a.latitudCliente,
									a.longitudCliente				
									FROM (
									SELECT codigousuario, asesor, codigocliente, cliente, tipovisita, fechaTarea, fechaVisita, comentario, imagen, latitud, longitud, latitudCliente, longitudCliente,clienteid
									from view_tareasdetalle
									where Date(fechaTarea) = $1 and idtipovisita = 1 and usuarioid=$2
									) a
									INNER JOIN (
									SELECT codigousuario, asesor, codigocliente, cliente, tipovisita, fechaTarea, fechaVisita, comentario, imagen, latitud, longitud, latitudCliente, longitudCliente,clienteid
									from view_tareasdetalle
									where  Date(fechaTarea) = $1 and idtipovisita = 2 and usuarioid= $2
									) b on Date(a.fechaTarea)  = Date(b.fechaTarea)  and a.clienteid =b.clienteid`, fecha, usuarioId)
	if err != nil {
		return []model.TareaHorasModelReporteWeb{}, err
	}
	defer rows.Close()
	tareas := []model.TareaHorasModelReporteWeb{}
	for rows.Next() {
		var tarea model.TareaHorasModelReporteWeb
		err := rows.Scan(&tarea.Codigousuario, &tarea.Respomsable, &tarea.CodigoCliente, &tarea.Cliente, &tarea.Fecha, &tarea.FechaEntrada, &tarea.FechaSalida, &tarea.ComentarioEntrada, &tarea.ComentarioSalida, &tarea.ImagenEntrada, &tarea.ImagenSalida, &tarea.UbicacionEntrada, &tarea.UbicacionSalida, &tarea.LatitudEntrada, &tarea.LongitudEntrada, &tarea.LatitudSalida, &tarea.LongitudSalida, &tarea.LatitudCliente, &tarea.LongitudCliente)
		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}
	return tareas, nil
}
