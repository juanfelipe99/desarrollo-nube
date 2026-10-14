package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Tiempo total para revisar todas las dependencias.
const tiempoMaxSalud = 3 * time.Second

const (
	estadoOK        = "ok"
	estadoError     = "error"
	estadoDegradado = "degradado"
)

type ComprobadorSalud func(ctx context.Context) error

type estadoSalud struct {
	Estado      string            `json:"estado"`
	Componentes map[string]string `json:"componentes"`
}

// salud responde 200 si todas las dependencias responden y 503 si alguna falla.
func salud(comprobadores map[string]ComprobadorSalud) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancelar := context.WithTimeout(c.Request.Context(), tiempoMaxSalud)
		defer cancelar()

		respuesta := estadoSalud{Estado: estadoOK, Componentes: map[string]string{}}
		codigo := http.StatusOK

		for nombre, comprobar := range comprobadores {
			resultado := estadoOK

			if err := comprobar(ctx); err != nil {
				resultado = estadoError
				respuesta.Estado = estadoDegradado
				codigo = http.StatusServiceUnavailable
			}

			respuesta.Componentes[nombre] = resultado
		}

		c.JSON(codigo, respuesta)
	}
}
