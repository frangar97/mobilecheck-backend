package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type VisitaRepository interface {
	CrearVisita(context.Context, model.CreateVisitaModel, string, int64) (int64, error)
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
