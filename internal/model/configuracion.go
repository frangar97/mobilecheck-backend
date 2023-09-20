package model

type ConfiguracionSubcidioModel struct {
	Id        int64  `json:"id"`
	Nombre    string `json:"nombre"`
	Parametro string `json:"parametro"`
	Maxlength int64  `json:"maxlength"`
	Minlength int64  `json:"minlength"`
}

type ConfiguracionSubcidioUpdateModel struct {
	Id        int64  `json:"id"`
	Parametro string `json:"parametro"`
}
