package config

import (
	"strings"
	"testing"
)

func TestCargarZonaHorariaRechazaZonaInvalida(t *testing.T) {
	t.Setenv(appZonaHoraria.nombre, "Marte/Monte_Olimpo")

	_, err := cargarZonaHoraria()

	if err == nil || !strings.Contains(err.Error(), appZonaHoraria.nombre) {
		t.Fatalf("se esperaba un error que mencionara %s: %v", appZonaHoraria.nombre, err)
	}
}
