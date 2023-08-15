package model

type PaisModel struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Abreviatura string `json:"apellido"`
	Activo      string `json:"telefono"`
}

type CreatePaisModel struct {
	Nombre      string `json:"nombre"`
	Abreviatura string `json:"apellido"`
	Activo      string `json:"telefono"`
}

type UpdatePaisModel struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Abreviatura string `json:"apellido"`
	Activo      string `json:"telefono"`
}

type EstadoPaisModel struct {
	ID     int64  `json:"id"`
	Activo string `json:"telefono"`
}
