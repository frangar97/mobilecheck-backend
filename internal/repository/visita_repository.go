package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type VisitaRepository interface {
	CrearVisita(context.Context, model.CreateVisitaModel, string, int64) (int64, error)
	ObtenerVisitasPorUsuario(context.Context, int64) ([]model.VisitaModel, error)
	ObtenerVisitasPorUsuarioDelDia(context.Context, string, int64) ([]model.VisitaModel, error)
	ObtenerVisitaPorId(context.Context, int64) (model.VisitaModel, error)
}

type visitaRepositoryImpl struct {
	db *sql.DB
}

func newVisitaRepository(db *sql.DB) *visitaRepositoryImpl {
	return &visitaRepositoryImpl{db: db}
}

func (v *visitaRepositoryImpl) CrearVisita(ctx context.Context, visita model.CreateVisitaModel, imagenUrl string, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := v.db.QueryRowContext(ctx, "INSERT INTO Visita(comentario,latitud,longitud,fecha,imagen,usuarioId,clienteId,tipoVisitaId) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id", visita.Comentario, visita.Latitud, visita.Longitud, visita.Fecha, imagenUrl, usuarioId, visita.ClienteId, visita.TipoVisitaId).Scan(&idGenerado)

	return idGenerado, err
}

func (v *visitaRepositoryImpl) ObtenerVisitasPorUsuario(ctx context.Context, usuarioId int64) ([]model.VisitaModel, error) {
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
		WHERE V.usuarioId = $1
		ORDER BY V.fecha DESC
	`, usuarioId)

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
		INNER JOIN TipoVisita TV ON V.tipoVisitaId = TV.id
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
	`, visitaId).Scan(&visita.ID, &visita.Comentario, &visita.Latitud, &visita.Longitud, &visita.Imagen, &visita.Fecha, &visita.Cliente, &visita.TipoVisita, &visita.Color)

	return visita, err
}
