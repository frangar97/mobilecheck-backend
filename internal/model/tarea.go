package model

import "time"

type TareaModelMovil struct {
	ID          int64     `json:"id"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	Completada  bool      `json:"completada"`
	ClienteId   int64     `json:"clienteId"`
	Cliente     string    `json:"cliente"`
}

type TareaModelWeb struct {
	ID          int64     `json:"id"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	Completada  bool      `json:"completada"`
}

type CreateTareaModelMovil struct {
	Descripcion string `json:"descripcion" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	ClienteId   int64  `json:"clienteId" binding:"required"`
}

type CreateTareaModelWeb struct {
	Descripcion string `json:"descripcion" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	ClienteId   int64  `json:"clienteId" binding:"required"`
	UsuarioId   int64  `json:"usuarioId" binding:"required"`
}

type CantidadTareaPorUsuario struct {
	Nombre      string `json:"nombre"`
	Completadas int    `json:"completadas"`
	Pendientes  int    `json:"pendientes"`
	Total       int    `json:"total"`
}
