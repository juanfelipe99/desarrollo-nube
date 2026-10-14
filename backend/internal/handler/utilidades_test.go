package handler

import (
	"context"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Utilidades compartidas por las pruebas del handler.

type entorno struct {
	router *gin.Engine
}

// configurar permite que cada prueba cambie las dependencias.
func nuevoEntorno(t *testing.T, configurar func(*Dependencias)) entorno {
	t.Helper()

	gin.SetMode(gin.TestMode)

	d := Dependencias{
		Salud: map[string]ComprobadorSalud{"postgres": func(context.Context) error { return nil }},
	}

	if configurar != nil {
		configurar(&d)
	}

	return entorno{router: NuevoRouter(d)}
}

func (e entorno) peticion(metodo, ruta, cuerpo string) *httptest.ResponseRecorder {
	var lector io.Reader
	if cuerpo != "" {
		lector = strings.NewReader(cuerpo)
	}

	req := httptest.NewRequest(metodo, ruta, lector)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)

	return w
}

func esperarEstado(t *testing.T, w *httptest.ResponseRecorder, estado int) {
	t.Helper()

	if w.Code != estado {
		t.Fatalf("estado = %d, se esperaba %d; cuerpo: %s", w.Code, estado, w.Body.String())
	}
}
