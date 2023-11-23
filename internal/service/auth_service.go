package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/frangar97/mobilecheck-backend/internal/model"
	"github.com/frangar97/mobilecheck-backend/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginWeb(context.Context, model.AuthCredencialModel) (string, []model.PantallaAccesoModel, error)
	LoginMovil(context.Context, model.AuthCredencialModel) (string, string, error)
}

type authServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
	accesoReposcitory repository.AccesoRepository
}

func newAuthService(usuarioRepository repository.UsuarioRepository, accesoReposcitory repository.AccesoRepository) *authServiceImpl {
	return &authServiceImpl{usuarioRepository: usuarioRepository, accesoReposcitory: accesoReposcitory}
}

func (a *authServiceImpl) LoginWeb(ctx context.Context, credenciales model.AuthCredencialModel) (string, []model.PantallaAccesoModel, error) {
	usuario, err := a.usuarioRepository.ObtenerPorUsuario(ctx, credenciales.Usuario)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, fmt.Errorf("el usuario %s no se encuentra registrado", credenciales.Usuario)
		}

		return "", nil, err
	}

	if !usuario.Activo {
		return "", nil, fmt.Errorf("el usuario no se encuentra activo")
	}

	if !usuario.Web {
		return "", nil, fmt.Errorf("el usuario no tiene acceso para la aplicación web")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password))

	if err != nil {
		return "", nil, fmt.Errorf("usuario o contraseña incorrecto")
	}

	pantallasAcceso, err := a.accesoReposcitory.ObtenerPantallasAccesos(ctx, usuario.ID, false, true)

	if err != nil {
		return "", nil, fmt.Errorf("Error al obtener los permisos del usuario")
	}

	// if len(pantallasAcceso) <= 0 {
	// 	return "", nil, fmt.Errorf("Accesos bloqueados")
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuarioId": usuario.ID,
		"web":       usuario.Web,
		"movil":     usuario.Movil,
		"paisId":    usuario.PaisId,
	})

	tokenString, err := token.SignedString([]byte("ProbandoTokenSeguridad"))

	if err != nil {
		return "", nil, err
	}

	return tokenString, pantallasAcceso, nil
}

func (a *authServiceImpl) LoginMovil(ctx context.Context, credenciales model.AuthCredencialModel) (string, string, error) {
	usuario, err := a.usuarioRepository.ObtenerPorUsuario(ctx, credenciales.Usuario)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("el usuario %s no se encuentra registrado", credenciales.Usuario)
		}

		return "", "", err
	}

	if !usuario.Activo {
		return "", "", fmt.Errorf("el usuario no se encuentra activo")
	}

	if !usuario.Movil {
		return "", "", fmt.Errorf("el usuario no tiene acceso para la aplicación movil")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password))

	if err != nil {
		return "", "", fmt.Errorf("usuario o contraseña incorrecto")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuarioId": usuario.ID,
		"web":       usuario.Web,
		"movil":     usuario.Movil,
	})

	tokenString, err := token.SignedString([]byte("ProbandoTokenSeguridad"))

	if err != nil {
		return "", "", err
	}

	return tokenString, fmt.Sprintf("%s %s", usuario.Nombre, usuario.Apellido), nil
}
