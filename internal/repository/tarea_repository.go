package repository

import (
	"context"
	"database/sql"

	"github.com/frangar97/mobilecheck-backend/internal/model"
)

type TareaRepository interface {
	CrearTarea(context.Context, model.CreateTareaModel) (int64, error)
}

type tareaRepositoryImpl struct {
	db *sql.DB
}

func newTareaRepository(db *sql.DB) *tareaRepositoryImpl {
	return &tareaRepositoryImpl{
		db: db,
	}
}

func (t *tareaRepositoryImpl) CrearTarea(ctx context.Context, tarea model.CreateTareaModel) (int64, error) {
	var idGenerado int64

	err := t.db.QueryRowContext(ctx, "INSERT INTO Tarea(descripcion,fecha,clienteId) VALUES($1,$2,$3) RETURNING id", tarea.Descripcion, tarea.Fecha, tarea.ClienteId).Scan(&idGenerado)

	return idGenerado, err
}
