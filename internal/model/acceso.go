package model

import "time"

// Accesos en pantalla
// ----------Modelos permisos Asignados
type MenuUsuarioModel struct {
	Opcion    string                 `json:"opcion"`
	Pantallas []PantallaUsuarioModel `json:"pantallas"`
}

type PantallaUsuarioModel struct {
	Pantalla string               `json:"pantalla"`
	Objetos  []ObjetoUsuarioModel `json:"objetos"`
}

type ObjetoUsuarioModel struct {
	Objeto string `json:"objeto"`
}

// ----------Modelos Obtienen permisos Asignados
type MenuUsuarioDataModel struct {
	ID     int64  `json:"id"`
	Opcion string `json:"opcion"`
}

type PantallaUsuarioDataModel struct {
	ID       int64  `json:"id"`
	Pantalla string `json:"pantalla"`
}

//-----Obtener , nuevo y editar permisos usuario

type MenuAsignadoUsuario struct {
	MenuAsignado   []MenuUsuarioDataModel `json:"menuAsignado"`
	MenuNoAsignado []MenuUsuarioDataModel `json:"menuNoAsignado"`
}

type MenuOpcionUsuarioModel struct {
	IdMenuOpcion int64 `json:"idMenuOpcion"`
	Activo       bool  `json:"activo"`
}

type PantallaAsignadaUsuario struct {
	PantallaAsignada   []PantallaUsuarioDataModel `json:"pantallaAsignada"`
	PantallaNoAsignada []PantallaUsuarioDataModel `json:"pantallaNoAsignada"`
}

type PantallanUsuarioModel struct {
	IdPantalla int64 `json:"idPantalla"`
	Activo     bool  `json:"activo"`
}

type AsignarMenuUsuarioModel struct {
	Id            int64     `json:"id"`
	Idusuario     int64     `json:"idusuario" binding:"required"`
	Idmenuopcion  int64     `json:"idmenuopcion" binding:"required"`
	Activo        *bool     `json:"activo" binding:"required"`
	UsuarioAccion int64     `json:"usuarioAccion"`
	FechaAccion   time.Time `json:"fechaAccion"`
}

type CreatePantallaUsuarioModel struct {
	Idusuario   int64     `json:"idusuario" binding:"required"`
	Idpantalla  int64     `json:"idpantalla" binding:"required"`
	Activo      *bool     `json:"activo" binding:"required"`
	UsuarioCrea int64     `json:"UsuarioCrea"`
	FechaCrea   time.Time `json:"fechacrea"`
}
