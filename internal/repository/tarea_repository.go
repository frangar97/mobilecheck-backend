package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TareaRepository interface {
	ObtenerTareaPorIdMovil(context.Context, int64) (model.TareaModelMovil, error)
	CrearTareaMovil(context.Context, model.CreateTareaModelMovil, int64) (int64, error)
	ObtenerTareasDelDia(context.Context, string, int64) ([]model.TareaModelMovil, error)
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

	err := t.db.QueryRowContext(ctx, "SELECT id,descripcion,fecha,completada FROM Tarea WHERE id = $1 LIMIT 1", tareaId).Scan(&tareaModel.ID, &tareaModel.Descripcion, &tareaModel.Fecha, &tareaModel.Completada)

	return tareaModel, err
}

func (t *tareaRepositoryImpl) CrearTareaMovil(ctx context.Context, tarea model.CreateTareaModelMovil, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO Tarea(descripcion,fecha,clienteId,usuarioId,completada) VALUES($1,$2,$3,$4,false) RETURNING id", tarea.Descripcion, tarea.Fecha, tarea.ClienteId, usuarioId).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tareaRepositoryImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	rows, err := t.db.QueryContext(ctx, "SELECT id,descripcion,fecha,completada FROM Tarea WHERE DATE(fecha) = $1 AND usuarioId = $2", fecha, usuarioId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tareas []model.TareaModelMovil

	for rows.Next() {
		var tarea model.TareaModelMovil

		err := rows.Scan(&tarea.ID, &tarea.Descripcion, &tarea.Fecha, &tarea.Completada)
		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	return tareas, nil
}
