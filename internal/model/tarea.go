package model

import (
	"mime/multipart"
	"time"
)

type TareaModelMovil struct {
	ID                   int64     `json:"id"`
	Meta                 string    `json:"meta"`
	Fecha                time.Time `json:"fecha"`
	Completada           bool      `json:"completada"`
	ClienteId            int64     `json:"clienteId"`
	Cliente              string    `json:"cliente"`
	ImagenRequerida      bool      `json:"imagenRequerida"`
	TipoVisita           string    `json:"tipoVisita"`
	Requieremeta         bool      `json:"requieremeta"`
	MetaLinea            string    `json:"metaLinea"`
	MetaSublinea         string    `json:"metaSublinea"`
	RequieremetaLinea    bool      `json:"requiereMetaLinea"`
	RequieremetaSubLinea bool      `json:"requiereMetaSubLinea"`
}

type TareaModelWeb struct {
	ID                   int64     `json:"id"`
	Fecha                time.Time `json:"fecha"`
	Completada           bool      `json:"completada"`
	Cliente              string    `json:"cliente"`
	ImagenRequerida      bool      `json:"imagenRequerida"`
	Asesor               string    `json:"asesor"`
	TipoVisitaId         int64     `json:"tipoVisitaId"`
	Meta                 string    `json:"meta"`
	Latitud              float64   `json:"latitud"`
	Longitud             float64   `json:"longitud"`
	Imagen               string    `json:"imagen"`
	TipoVisita           string    `json:"tipoVisita"`
	Comentario           string    `json:"comentario"`
	Requieremeta         bool      `json:"requiereMeta"`
	MetaAsignada         string    `json:"metaAsignada"`
	MetaCumplida         string    `json:"metaCumplida"`
	CodigoUsuario        string    `json:"codigoUsuario"`
	UsuarioId            int64     `json:"usuarioId"`
	MetaLineaAsignada    string    `json:"metaLineaAsignada"`
	MetaSubLineaAsignada string    `json:"metaSubLineaAsignada"`
	MetaLineaCumplida    string    `json:"metaLineaCumplida"`
	MetaSubLineaCumplida string    `json:"metaSubLineaCumplida"`
	CodigoCliente        string    `json:"codigoCliente"`
	LatitudCliente       float64   `json:"latitudCliente"`
	LongitudCliente      float64   `json:"longitudCliente"`
}

type CreateTareaModelMovil struct {
	Descripcion string `json:"descripcion" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	ClienteId   int64  `json:"clienteId" binding:"required"`
}

type CreateTareaModelWeb struct {
	Fecha           string    `json:"fecha" binding:"required"`
	ClienteId       int64     `json:"clienteId" binding:"required"`
	UsuarioId       int64     `json:"usuarioId" binding:"required"`
	TipoVisitaId    int64     `json:"tipoVisitaId" binding:"required"`
	Meta            string    `json:"meta"`
	ImagenRequerida *bool     `json:"imagenRequerida" binding:"required"`
	MetaLinea       string    `json:"metaLinea"`
	MetaSubLinea    string    `json:"metaSubLinea"`
	UsuarioCrea     int64     `json:"UsuarioCrea"`
	FechaCrea       time.Time `json:"fechacrea"`
}

type CantidadTareaPorUsuario struct {
	Nombre      string `json:"nombre"`
	Completadas int    `json:"completadas"`
	Pendientes  int    `json:"pendientes"`
	Total       int    `json:"total"`
}

type CompletarTareaModel struct {
	Comentario      string                `form:"comentario"`
	Latitud         float64               `form:"latitud" binding:"required"`
	Longitud        float64               `form:"longitud" binding:"required"`
	Fecha           string                `form:"fecha" binding:"required"`
	Imagen          *multipart.FileHeader `form:"imagen"`
	ClienteId       int64                 `form:"clienteId" binding:"required"`
	Meta            string                `form:"meta" `
	TareaId         int64                 `form:"tareaId" binding:"required"`
	ImagenRequerida *bool                 `form:"imagenRequerida" binding:"required"`
	MetaLinea       string                `form:"metaLinea" `
	MetaSubLinea    string                `form:"metaSubLinea" `
}

type CreateTareaMasivaModelWeb struct {
	Meta            string   `json:"meta"`
	Fechas          []string `json:"fechas" binding:"required"`
	ClienteId       int64    `json:"clienteId" binding:"required"`
	UsuarioId       int64    `json:"usuarioId" binding:"required"`
	ImagenRequerida *bool    `json:"imagenRequerida" binding:"required"`
	TipoVisitaId    int64    `json:"tipoVisitaId" binding:"required"`
	MetaLinea       string   `json:"metaLinea"`
	MetaSubLinea    string   `json:"metaSubLinea"`
}

type CreateTareasExcelWeb struct {
	Tareas []CreateTareaModelWeb `json:"tareas" binding:"required"`
}

type ValidarTareasExcelWeb struct {
	Cliente     string `json:"cliente"`
	Responsable string `json:"responsable"`
	TipoVisita  string `json:"tipoVisita"`
	Tarea       string `json:"tarea"`
	Error       bool   `json:"error"`
}

type TareaHorasModelReporteWeb struct {
	Codigousuario     string    `json:"codigoUsuario"`
	Respomsable       string    `json:"respomsable"`
	CodigoCliente     string    `json:"codigoCliente"`
	Cliente           string    `json:"cliente"`
	Fecha             string    `json:"fecha"`
	FechaEntrada      time.Time `json:"fechaEntrada"`
	FechaSalida       time.Time `json:"fechaSalida"`
	HorasTrabajadas   string    `json:"horasTrabajadas"`
	ComentarioEntrada string    `json:"comentarioEntrada"`
	ComentarioSalida  string    `json:"comentarioSalida"`
	ImagenEntrada     string    `json:"imagenEntrada"`
	ImagenSalida      string    `json:"imagenSalida"`
	UbicacionEntrada  string    `json:"ubicacionEntrada"`
	UbicacionSalida   string    `json:"ubicacionSalida"`
	LatitudEntrada    float64   `json:"latitudEntrada"`
	LongitudEntrada   float64   `json:"longitudEntrada"`
	LatitudSalida     float64   `json:"latitudSalida"`
	LongitudSalida    float64   `json:"longitudSalida"`
	LatitudCliente    float64   `json:"latitudCliente"`
	LongitudCliente   float64   `json:"longitudCliente"`
}

type ParamReportTareasHoras struct {
	UsuarioId []int    `json:"usuarioId"`
	Fechas    []string `json:"fechas"`
}
