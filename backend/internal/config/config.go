// Package config carga la configuración desde variables de entorno.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Entorno     string // desarrollo | produccion
	PuertoHTTP  string
	ZonaHoraria *time.Location
}

func (c Config) EsDesarrollo() bool { return c.Entorno == "desarrollo" }

func (c Config) EsProduccion() bool { return c.Entorno == "produccion" }

// Cargar reúne todos los errores para mostrarlos juntos al arrancar.
func Cargar() (Config, error) {
	var errs []error

	c := Config{
		Entorno:    leerVariable("APP_ENTORNO", "desarrollo"),
		PuertoHTTP: leerVariable("APP_PUERTO", "8080"),
	}

	zona, err := time.LoadLocation(leerVariable("APP_ZONA_HORARIA", "America/Bogota"))
	if err != nil {
		errs = append(errs, fmt.Errorf("APP_ZONA_HORARIA: %w", err))
	}
	c.ZonaHoraria = zona

	return c, errors.Join(errs...)
}

func leerVariable(variable, porDefecto string) string {
	if v := strings.TrimSpace(os.Getenv(variable)); v != "" {
		return v
	}

	return porDefecto
}
