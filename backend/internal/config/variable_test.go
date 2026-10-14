package config

import (
	"strings"
	"testing"
)

func TestLeerQuitaEspacios(t *testing.T) {
	v := opcional("PRUEBA_VARIABLE", "por-defecto")
	t.Setenv(v.nombre, "  valor  ")

	valor, err := v.leer()

	if err != nil || valor != "valor" {
		t.Errorf("leer() = %q, %v, se esperaba %q", valor, err, "valor")
	}
}

func TestLeerObligatoriaVaciaDevuelveError(t *testing.T) {
	v := obligatoria("PRUEBA_VARIABLE")
	t.Setenv(v.nombre, "")

	_, err := v.leer()

	if err == nil || !strings.Contains(err.Error(), v.nombre) {
		t.Errorf("se esperaba un error que mencionara %s: %v", v.nombre, err)
	}
}

func TestLeerOpcionalVaciaDevuelvePorDefecto(t *testing.T) {
	v := opcional("PRUEBA_VARIABLE", "por-defecto")
	t.Setenv(v.nombre, "")

	valor, err := v.leer()

	if err != nil || valor != v.porDefecto {
		t.Errorf("leer() = %q, %v, se esperaba %q", valor, err, v.porDefecto)
	}
}
