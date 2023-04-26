package model

type TipoVisitaModel struct {
	ID           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	Color        string `json:"color"`
	Activo       bool   `json:"activo"`
	RequiereMeta bool   `json:"requiereMeta"`
}

type CreateTipoVisitaModel struct {
	Nombre       string `json:"nombre" binding:"required"`
	Color        string `json:"color" binding:"required"`
	RequiereMeta *bool  `json:"requiereMeta" binding:"required"`
}

type UpdateTipoVisitaModel struct {
	Nombre       string `json:"nombre" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Activo       *bool  `json:"activo" binding:"required"`
	RequiereMeta *bool  `json:"requiereMeta" binding:"required"`
}
