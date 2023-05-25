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
	Meta       string    `json:"meta"`
	Color      string    `json:"color"`
	Fecha      time.Time `json:"fecha"`
	TipoVisita string    `json:"tipoVisita"`
}

type CreateVisitaModel struct {
	Comentario   string                `form:"comentario" binding:"required"`
	Latitud      float64               `form:"latitud" binding:"required"`
	Longitud     float64               `form:"longitud" binding:"required"`
	Fecha        string                `form:"fecha" binding:"required"`
	Imagen       *multipart.FileHeader `form:"imagen" binding:"required"`
	ClienteId    int64                 `form:"clienteId" binding:"required"`
	Meta         string                `form:"meta"`
	MetaLinea    string                `form:"metaLinea" `
	MetaSubLinea string                `form:"metaSubLinea" `
}

type CantidadVisitaPorUsuario struct {
	Nombre   string `json:"nombre"`
	Cantidad int    `json:"cantidad"`
}

type CantidadVisitaPorTipo struct {
	Nombre   string `json:"nombre"`
	Color    string `json:"color"`
	Cantidad int    `json:"cantidad"`
}

type VisitaTareaModel struct {
	ID                   int64     `json:"id"`
	Cliente              string    `json:"cliente"`
	Comentario           string    `json:"comentario"`
	Latitud              float64   `json:"latitud"`
	Longitud             float64   `json:"longitud"`
	Imagen               string    `json:"imagen"`
	TipoVisita           string    `json:"tipoVisita"`
	Fecha                time.Time `json:"fecha"`
	MetaVisita           string    `json:"metaVisita"`
	MetaTarea            string    `json:"metaTarea"`
	RequiereMeta         bool      `json:"requiereMeta"`
	MetaLineaTarea       string    `json:"metaLineaTarea"`
	MetaLineaVisita      string    `json:"metaLineaVisita"`
	MetaSubLineaTarea    string    `json:"metaSubLineaTarea"`
	MetaSubLineaVisita   string    `json:"metaSubLineaVisita"`
	LatitudCliente       float64   `json:"latitudCliente"`
	LongitudCliente      float64   `json:"longitudCliente"`
	RequiereMetaLinea    bool      `json:"requiereMetaLinea"`
	RequiereMetaSubLinea bool      `json:"requiereMetaSubLinea"`
}
