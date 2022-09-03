package model

import (
	"mime/multipart"
	"time"
)

type CreateVisitaModel struct {
	Comentario   string                `form:"comentario" binding:"required"`
	Latitud      float64               `form:"latitud" binding:"required"`
	Longitud     float64               `form:"longitud" binding:"required"`
	Fecha        time.Time             `form:"fecha" binding:"required"`
	Imagen       *multipart.FileHeader `form:"imagen" binding:"required"`
	ClienteId    int64                 `form:"clienteId" binding:"required"`
	TipoVisitaId int64                 `form:"tipoVisitaId" binding:"required"`
}
