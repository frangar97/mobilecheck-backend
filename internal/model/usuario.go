package model

type UsuarioModel struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	Activo   bool   `json:"activo"`
	Usuario  string `json:"usuario"`
	Password string `json:"-"`
	Web      bool   `json:"web"`
	Movil    bool   `json:"movil"`
	PaisId   int64  `json:"paisid"`
}

type CreateUsuarioModel struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	Usuario  string `json:"usuario" binding:"required"`
	Password string `json:"password" binding:"required"`
	Web      *bool  `json:"web" binding:"required"`
	Movil    *bool  `json:"movil" binding:"required"`
	PaisId   int64  `json:"paisid" binding:"required"`
}

type UpdateUsuarioModel struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	Activo   *bool  `json:"activo" binding:"required"`
	Usuario  string `json:"usuario" binding:"required"`
	Web      *bool  `json:"web" binding:"required"`
	Movil    *bool  `json:"movil" binding:"required"`
	PaisId   int64  `json:"paisid" binding:"required"`
}

type UpdatePasswordModel struct {
	Id       int64  `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CodigoUsuarioModel struct {
	CodigoUsuario string `json:"CodigoUsuario"`
	TipoContrato  string `json:"tipoContrato"`
	Nombre        string `json:"nombre"`
}

type ImpulsadorasPayRollModel struct {
	Codigo       string `json:"codigo"`
	Nombre       string `json:"nombre"`
	NumeroCuenta string `json:"numeroCuenta"`
	TipoContrato string `json:"tipoContrato"`
	Valor        string `json:"valor"`
	Estado       string `json:"estado"`
}

type ImpulsadorasPayRollDataModel struct {
	SinCodigoPayRoll       []ImpulsadorasPayRollModel `json:"sinCodigoPayRoll"`
	Reporte                []ImpulsadorasPayRollModel `json:"reporte"`
	NoEncontradosEnPayroll []ImpulsadorasPayRollModel `json:"noEncontrados"`
}
