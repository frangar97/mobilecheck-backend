package model

type UsuarioModel struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	Activo   bool   `json:"activo"`
	Usuario  string `json:"usuario"`
	Web      bool   `json:"web"`
	Movil    bool   `json:"movil"`
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
}
