package model

import (
	"mime/multipart"
	"time"
)

type TareaModelMovil struct {
	ID              int64     `json:"id"`
	Meta            string    `json:"meta"`
	Fecha           time.Time `json:"fecha"`
	Completada      bool      `json:"completada"`
	ClienteId       int64     `json:"clienteId"`
	Cliente         string    `json:"cliente"`
	ImagenRequerida bool      `json:"imagenRequerida"`
	TipoVisita      string    `json:"tipoVisita"`
	Requieremeta    bool      `json:"requieremeta"`
}

type TareaModelWeb struct {
	ID              int64     `json:"id"`
	Fecha           time.Time `json:"fecha"`
	Completada      bool      `json:"completada"`
	Cliente         string    `json:"cliente"`
	ImagenRequerida bool      `json:"imagenRequerida"`
	Asesor          string    `json:"asesor"`
	TipoVisitaId    int64     `json:"tipoVisitaId"`
	Meta            string    `json:"meta"`
	Latitud         float64   `json:"latitud"`
	Longitud        float64   `json:"longitud"`
	Imagen          string    `json:"imagen"`
	TipoVisita      string    `json:"tipoVisita"`
	Comentario      string    `json:"comentario"`
	Requieremeta    bool      `json:"requiereMeta"`
	MetaAsignada    string    `json:"metaAsignada"`
	MetaCumplida    string    `json:"metaCumplida"`
	CodigoUsuario   string    `json:"codigoUsuario"`
	UsuarioId       int64     `json:"usuarioId"`
}

type CreateTareaModelMovil struct {
	Descripcion string `json:"descripcion" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	ClienteId   int64  `json:"clienteId" binding:"required"`
}

type CreateTareaModelWeb struct {
	Fecha           string `json:"fecha" binding:"required"`
	ClienteId       int64  `json:"clienteId" binding:"required"`
	UsuarioId       int64  `json:"usuarioId" binding:"required"`
	TipoVisitaId    int64  `json:"tipoVisitaId" binding:"required"`
	Meta            string `json:"meta"`
	ImagenRequerida *bool  `json:"imagenRequerida" binding:"required"`
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
}

type CreateTareaMasivaModelWeb struct {
	Meta            string   `json:"meta"`
	Fechas          []string `json:"fechas" binding:"required"`
	ClienteId       int64    `json:"clienteId" binding:"required"`
	UsuarioId       int64    `json:"usuarioId" binding:"required"`
	ImagenRequerida *bool    `json:"imagenRequerida" binding:"required"`
	TipoVisitaId    int64    `json:"tipoVisitaId" binding:"required"`
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
	Codigousuario     string `json:"codigousuario"`
	Cliente           string `json:"cliente"`
	Asesor            string `json:"asesor"`
	EntradaCompletada bool   `json:"EntradaCompletada"`
	SalidaCompletada  bool   `json:"salidaCompletada"`
	Fecha             string `json:"fecha"`
	HoraEntrada       string `json:"horaEntrada"`
	HoraSalida        string `json:"horaSalida"`
	HorasTrabajadas   string `json:"horasTrabajadas"`
	ComentarioEntrada string `json:"comentarioEntrada"`
	ComentarioSalida  string `json:"comentarioSalida"`
	UbicacionEntrada  string `json:"ubicacionEntrada"`
	UbicacionSalida   string `json:"ubicacionSalida"`
	ImagenEntrada     string `json:"imagenEntrada"`
	ImaenSalida       string `json:"imaenSalida"`
}

type ParamReportTareasHoras struct {
	UsuarioId []int    `json:"usuarioId"`
	Fechas    []string `json:"fechas"`
}
