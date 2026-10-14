package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
)

func TestSaludResponde200(t *testing.T) {
	e := nuevoEntorno(t, nil)

	w := e.peticion(http.MethodGet, "/salud", "")

	esperarEstado(t, w, http.StatusOK)

	r := leerEstadoSalud(t, w.Body.Bytes())
	if r.Estado != estadoOK || r.Componentes["postgres"] != estadoOK {
		t.Errorf("respuesta inesperada: %+v", r)
	}
}

func TestSaludResponde503SiUnaDependenciaFalla(t *testing.T) {
	e := nuevoEntorno(t, func(d *Dependencias) {
		d.Salud["redis"] = func(context.Context) error { return errors.New("caído") }
	})

	w := e.peticion(http.MethodGet, "/salud", "")

	esperarEstado(t, w, http.StatusServiceUnavailable)

	r := leerEstadoSalud(t, w.Body.Bytes())
	if r.Estado != estadoDegradado || r.Componentes["redis"] != estadoError || r.Componentes["postgres"] != estadoOK {
		t.Errorf("respuesta inesperada: %+v", r)
	}
}

func leerEstadoSalud(t *testing.T, cuerpo []byte) estadoSalud {
	t.Helper()

	var r estadoSalud
	if err := json.Unmarshal(cuerpo, &r); err != nil {
		t.Fatalf("la respuesta no es JSON válido: %v; cuerpo: %s", err, cuerpo)
	}

	return r
}
