package model

type AuthCredencialModel struct {
	Usuario  string `json:"usuario" binding:"required"`
	Password string `json:"password" binding:"required"`
}
