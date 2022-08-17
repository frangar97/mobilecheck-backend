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
