package model

import "time"

type ModuloModel struct {
	Id     int64  `json:"id"`
	Modulo string `json:"modulo"`
	Activo *bool  `json:"activo"`
}
type PantallaModel struct {
	Id       int64  `json:"id"`
	Pantalla string `json:"pantalla"`
	Web      *bool  `json:"web"`
	Movil    *bool  `json:"movil"`
	IdModulo int64  `json:"idModulo"`
}

type ObjetoModel struct {
	Id         int64  `json:"id"`
	IdPantalla int64  `json:"idPantalla"`
	Objeto     string `json:"objeto"`
}

type AccesoObjetollaModel struct {
	Id               int64 `json:"id"`
	IdObjeto         int64 `json:"idObjeto"`
	IdAccesoPantalla int64 `json:"idAccesoPantalla"`
	Activo           *bool `json:"activo"`
}

// type AsignarMenuUsuarioModel struct {
// 	Id            int64     `json:"id"`
// 	Idusuario     int64     `json:"idusuario" binding:"required"`
// 	Idmenuopcion  int64     `json:"idmenuopcion" binding:"required"`
// 	Activo        *bool     `json:"activo" binding:"required"`
// 	UsuarioAccion int64     `json:"usuarioAccion"`
// 	FechaAccion   time.Time `json:"fechaAccion"`
// }

// ----------- Acceso Pantalla -----
type CreateUpdateAccesoPantallaModel struct {
	IdPantalla      int64     `json:"idPantalla" binding:"required"`
	IdUsuario       int64     `json:"idUsuario" binding:"required"`
	Activo          bool      `json:"activo"`
	UsuarioCrea     int64     `json:"UsuarioCrea"`
	FechaCrea       time.Time `json:"fechacrea"`
	UsuarioModifica int64     `json:"usuariomodifica"`
	FechaModifica   time.Time `json:"fechamodifica"`
}

type AccesoPantallaModel struct {
	IdPantalla int64  `json:"idPantalla"`
	Pantalla   string `json:"pantalla"`
	Activo     bool   `json:"activo"`
	Web        bool   `json:"web"`
	Movil      bool   `json:"movil"`
}

type AccesoPantallaUsuarioModel struct {
	PantallasAsignadas   []AccesoPantallaModel `json:"pantallasAsignadas"`
	PantallasNoAsignadas []AccesoPantallaModel `json:"pantallasNoAsignadas"`
}

//---------------------------------

//-----------Accesos--------------
type PantallaAccesoModel struct {
	Pantalla string `json:"pantalla"`
}
