package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Tiempo total para revisar todas las dependencias.
const tiempoMaxSalud = 3 * time.Second

type ComprobadorSalud func(ctx context.Context) error

// Estado: "ok" o "degradado". Cada componente: "ok" o "error".
type estadoSalud struct {
	Estado      string            `json:"estado"`
	Componentes map[string]string `json:"componentes"`
}

// salud responde 200 si todas las dependencias responden y 503 si alguna falla.
func salud(comprobadores map[string]ComprobadorSalud) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancelar := context.WithTimeout(c.Request.Context(), tiempoMaxSalud)
		defer cancelar()

		respuesta := estadoSalud{Estado: "ok", Componentes: map[string]string{}}
		codigo := http.StatusOK

		for nombre, comprobar := range comprobadores {
			resultado := "ok"

			if err := comprobar(ctx); err != nil {
				resultado = "error"
				respuesta.Estado = "degradado"
				codigo = http.StatusServiceUnavailable
			}

			respuesta.Componentes[nombre] = resultado
		}

		c.JSON(codigo, respuesta)
	}
}
