package model

type ClienteModel struct {
	ID            int64   `json:"id"`
	CodigoCliente string  `json:"codigoCliente"`
	Nombre        string  `json:"nombre" `
	Telefono      string  `json:"telefono"`
	Email         string  `json:"email"`
	Direccion     string  `json:"direccion"`
	Latitud       float64 `json:"latitud" `
	Longitud      float64 `json:"longitud" `
	Usuario       string  `json:"usuario" `
	Activo        bool    `json:"activo"`
	UsuarioId     int64   `json:"usuarioId"`
}

type CreateClienteModel struct {
	CodigoCliente string  `json:"codigoCliente"`
	Nombre        string  `json:"nombre" binding:"required"`
	Telefono      string  `json:"telefono"`
	Email         string  `json:"email"`
	Direccion     string  `json:"direccion"`
	Latitud       float64 `json:"latitud"`
	Longitud      float64 `json:"longitud"`
	Activo        bool    `json:"activo"`
}

type UpdateClienteModel struct {
	CodigoCliente string  `json:"codigoCliente"`
	Nombre        string  `json:"nombre" binding:"required"`
	Telefono      string  `json:"telefono"`
	Email         string  `json:"email"`
	Direccion     string  `json:"direccion"`
	Latitud       float64 `json:"latitud"`
	Longitud      float64 `json:"longitud"`
	Activo        bool    `json:"activo"`
}
