package model

import (
	"mime/multipart"
	"time"
)

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
	Cliente     string    `json:"cliente"`
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

type CompletarTareaModel struct {
	Comentario   string                `form:"comentario" binding:"required"`
	Latitud      float64               `form:"latitud" binding:"required"`
	Longitud     float64               `form:"longitud" binding:"required"`
	Fecha        string                `form:"fecha" binding:"required"`
	Imagen       *multipart.FileHeader `form:"imagen" binding:"required"`
	ClienteId    int64                 `form:"clienteId" binding:"required"`
	TipoVisitaId int64                 `form:"tipoVisitaId" binding:"required"`
	TareaId      int64                 `form:"tareaId" binding:"required"`
}
