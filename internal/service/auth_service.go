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
	LoginWeb(context.Context, model.AuthCredencialModel) (string, error)
	LoginMovil(context.Context, model.AuthCredencialModel) (string, string, error)
}

type authServiceImpl struct {
	usuarioRepository repository.UsuarioRepository
}

func newAuthService(usuarioRepository repository.UsuarioRepository) *authServiceImpl {
	return &authServiceImpl{usuarioRepository: usuarioRepository}
}

func (a *authServiceImpl) LoginWeb(ctx context.Context, credenciales model.AuthCredencialModel) (string, error) {
	usuario, err := a.usuarioRepository.ObtenerPorUsuario(ctx, credenciales.Usuario)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("el usuario %s no se encuentra registrado", credenciales.Usuario)
		}

		return "", err
	}

	if !usuario.Activo {
		return "", fmt.Errorf("el usuario no se encuentra activo")
	}

	if !usuario.Web {
		return "", fmt.Errorf("el usuario no tiene acceso para la aplicación web")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password))

	if err != nil {
		return "", fmt.Errorf("usuario o contraseña incorrecto")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuarioId": usuario.ID,
		"web":       usuario.Web,
		"movil":     usuario.Movil,
		"paisId":    usuario.PaisId,
	})

	tokenString, err := token.SignedString([]byte("ProbandoTokenSeguridad"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
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
