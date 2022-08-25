package model

type ClienteModel struct {
	ID        int64   `json:"id"`
	Nombre    string  `json:"nombre" `
	Telefono  string  `json:"telefono"`
	Email     string  `json:"email"`
	Direccion string  `json:"direccion"`
	Latitud   float64 `json:"latitud" `
	Longitud  float64 `json:"longitud" `
	Usuario   string  `json:"usuario" `
}

type CreateClienteModel struct {
	Nombre    string  `json:"nombre" binding:"required"`
	Telefono  string  `json:"telefono"`
	Email     string  `json:"email"`
	Direccion string  `json:"direccion"`
	Latitud   float64 `json:"latitud" binding:"required"`
	Longitud  float64 `json:"longitud" binding:"required"`
	UsuarioID int64   `json:"usuarioId" binding:"required"`
}
