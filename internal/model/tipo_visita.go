package model

type TipoVisitaModel struct {
	ID                   int64  `json:"id"`
	Nombre               string `json:"nombre"`
	Color                string `json:"color"`
	Activo               bool   `json:"activo"`
	RequiereMeta         bool   `json:"requiereMeta"`
	RequiereMetaLinea    bool   `json:"requiereMetaLinea"`
	RequiereMetaSubLinea bool   `json:"requiereMetaSubLinea"`
}

type CreateTipoVisitaModel struct {
	Nombre               string `json:"nombre" binding:"required"`
	Color                string `json:"color" binding:"required"`
	RequiereMeta         *bool  `json:"requiereMeta" binding:"required"`
	RequiereMetaLinea    *bool  `json:"requiereMetaLinea" binding:"required"`
	RequiereMetaSubLinea *bool  `json:"requiereMetaSubLinea" binding:"required"`
}

type UpdateTipoVisitaModel struct {
	Nombre               string `json:"nombre" binding:"required"`
	Color                string `json:"color" binding:"required"`
	Activo               *bool  `json:"activo" binding:"required"`
	RequiereMeta         *bool  `json:"requiereMeta" binding:"required"`
	RequiereMetaLinea    *bool  `json:"requiereMetaLinea" binding:"required"`
	RequiereMetaSubLinea *bool  `json:"requiereMetaSubLinea" binding:"required"`
}
