package model

import "time"

type ConfiguracionSubcidioModel struct {
	Id        int64  `json:"id"`
	Nombre    string `json:"nombre"`
	Parametro string `json:"parametro"`
	Maxlength int64  `json:"maxlength"`
	Minlength int64  `json:"minlength"`
}

type ConfiguracionSubcidioUpdateModel struct {
	Id              int64     `json:"id" binding:"required"`
	Parametro       string    `json:"parametro" binding:"required"`
	UsuarioModifica int64     `json:"usuarioModifica" binding:"required"`
	FechaModifica   time.Time `json:"fechaModifica" binding:"required"`
}
