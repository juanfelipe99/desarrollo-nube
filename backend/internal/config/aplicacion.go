package config

import (
	"fmt"
	"time"
)

var (
	appEntorno     = opcional("APP_ENTORNO", "desarrollo")
	appPuerto      = opcional("APP_PUERTO", "8080")
	appZonaHoraria = opcional("APP_ZONA_HORARIA", "America/Bogota")
)

func cargarZonaHoraria() (*time.Location, error) {
	nombre, err := appZonaHoraria.leer()
	if err != nil {
		return nil, err
	}

	zona, err := time.LoadLocation(nombre)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", appZonaHoraria.nombre, err)
	}

	return zona, nil
}
