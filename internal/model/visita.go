package model

import (
	"mime/multipart"
	"time"
)

type VisitaModel struct {
	ID         int64     `json:"id"`
	Comentario string    `json:"comentario"`
	Latitud    float64   `json:"latitud"`
	Longitud   float64   `json:"longitud"`
	Imagen     string    `json:"imagen"`
	Cliente    string    `json:"cliente"`
	TipoVisita string    `json:"tipoVisita"`
	Color      string    `json:"color"`
	Fecha      time.Time `json:"fecha"`
}

type CreateVisitaModel struct {
	Comentario   string                `form:"comentario" binding:"required"`
	Latitud      float64               `form:"latitud" binding:"required"`
	Longitud     float64               `form:"longitud" binding:"required"`
	Fecha        string                `form:"fecha" binding:"required"`
	Imagen       *multipart.FileHeader `form:"imagen" binding:"required"`
	ClienteId    int64                 `form:"clienteId" binding:"required"`
	TipoVisitaId int64                 `form:"tipoVisitaId" binding:"required"`
}
