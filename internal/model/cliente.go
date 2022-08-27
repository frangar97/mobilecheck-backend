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
	Activo    bool    `json:"activo"`
}

type CreateClienteModel struct {
	Nombre    string  `json:"nombre" binding:"required"`
	Telefono  string  `json:"telefono"`
	Email     string  `json:"email"`
	Direccion string  `json:"direccion"`
	Latitud   float64 `json:"latitud" binding:"required"`
	Longitud  float64 `json:"longitud" binding:"required"`
}

type UpdateClienteModel struct {
	ID        int64   `json:"id" binding:"required"`
	Nombre    string  `json:"nombre" binding:"required"`
	Telefono  string  `json:"telefono" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Direccion string  `json:"direccion" binding:"required"`
	Latitud   float64 `json:"latitud" binding:"required"`
	Longitud  float64 `json:"longitud" binding:"required"`
	UsuarioId int64   `json:"usuarioId" binding:"required"`
	Activo    *bool   `json:"activo" binding:"required"`
}
