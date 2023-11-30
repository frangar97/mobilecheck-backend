package v1

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) webIdentity(c *gin.Context) {
	id, web, _, paisId, err := h.parseAuthHeaderWeb(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if !web {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no tiene acceso a la aplicación web"})
		return
	}

	c.Set("usuarioId", id)
	c.Set("paisId", paisId)
}

func (h *Handler) movilIdentity(c *gin.Context) {
	id, _, movil, err := h.parseAuthHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if !movil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no tiene acceso a la aplicación web"})
		return
	}

	VersionApp := c.GetHeader("VersionApp")
	versionActual := os.Getenv("VERSRION_MOVIL_APP")

	if VersionApp != versionActual {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "App desactualizada. Porfavor actualizar"})
		return
	}

	c.Set("usuarioId", id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (int64, bool, bool, error) {
	header := c.GetHeader("Authorization")

	if header == "" {
		return 0, false, false, errors.New("header de autorización vacio")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, false, false, errors.New("header de autorización invalido")
	}

	if len(headerParts[1]) == 0 {
		return 0, false, false, errors.New("token esta vacio")
	}

	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ProbandoTokenSeguridad"), nil
	})

	if err != nil {
		return 0, false, false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId := claims["usuarioId"]
		web := claims["web"]
		movil := claims["movil"]

		return int64(usuarioId.(float64)), web.(bool), movil.(bool), nil
	}

	return 0, false, false, errors.New("no se pudo obtener la identidad del usuario")
}

func (h *Handler) parseAuthHeaderWeb(c *gin.Context) (int64, bool, bool, int64, error) {
	header := c.GetHeader("Authorization")

	if header == "" {
		return 0, false, false, 0, errors.New("header de autorización vacio")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, false, false, 0, errors.New("header de autorización invalido")
	}

	if len(headerParts[1]) == 0 {
		return 0, false, false, 0, errors.New("token esta vacio")
	}

	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ProbandoTokenSeguridad"), nil
	})

	if err != nil {
		return 0, false, false, 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId := claims["usuarioId"]
		web := claims["web"]
		movil := claims["movil"]
		paisId := claims["paisId"]

		return int64(usuarioId.(float64)), web.(bool), movil.(bool), int64(paisId.(float64)), nil
	}

	return 0, false, false, 0, errors.New("no se pudo obtener la identidad del usuario")
}
