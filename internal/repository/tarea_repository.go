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

	err := t.db.QueryRowContext(ctx, "SELECT T.id,T.descripcion,T.fecha,T.completada,C.id,C.nombre FROM Tarea T INNER JOIN Cliente C ON T.clienteId = C.id WHERE T.id = $1 LIMIT 1", tareaId).Scan(&tareaModel.ID, &tareaModel.Descripcion, &tareaModel.Fecha, &tareaModel.Completada, &tareaModel.ClienteId, &tareaModel.Cliente)

	return tareaModel, err
}

func (t *tareaRepositoryImpl) CrearTareaMovil(ctx context.Context, tarea model.CreateTareaModelMovil, usuarioId int64) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO Tarea(descripcion,fecha,clienteId,usuarioId,completada) VALUES($1,$2,$3,$4,false) RETURNING id", tarea.Descripcion, tarea.Fecha, tarea.ClienteId, usuarioId).Scan(&idGenerado)

	return idGenerado, err
}

func (t *tareaRepositoryImpl) ObtenerTareasDelDia(ctx context.Context, fecha string, usuarioId int64) ([]model.TareaModelMovil, error) {
	rows, err := t.db.QueryContext(ctx, "SELECT T.id,T.descripcion,T.fecha,T.completada,C.id,C.nombre FROM Tarea T INNER JOIN Cliente C ON T.clienteId = C.id WHERE DATE(T.fecha) = $1 AND T.usuarioId = $2 ORDER BY T.fecha", fecha, usuarioId)
	if err != nil {
		return []model.TareaModelMovil{}, err
	}

	defer rows.Close()

	tareas := []model.TareaModelMovil{}

	for rows.Next() {
		var tarea model.TareaModelMovil

		err := rows.Scan(&tarea.ID, &tarea.Descripcion, &tarea.Fecha, &tarea.Completada, &tarea.ClienteId, &tarea.Cliente)
		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	return tareas, nil
}
