package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) webIdentity(c *gin.Context) {
	id, web, _, err := h.parseAuthHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if !web {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no tiene acceso a la aplicación web"})
		return
	}

	c.Set("usuarioId", id)
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
