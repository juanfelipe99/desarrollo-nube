package config

import (
	"errors"
	"time"
)

type Config struct {
	Entorno     string // desarrollo | produccion
	PuertoHTTP  string
	BaseDatos   string // URL de conexión
	ZonaHoraria *time.Location
}

func (c Config) EsDesarrollo() bool { return c.Entorno == "desarrollo" }

func (c Config) EsProduccion() bool { return c.Entorno == "produccion" }

// Cargar reúne todos los errores para mostrarlos juntos al arrancar.
func Cargar() (Config, error) {
	entorno, errEntorno := appEntorno.leer()
	puerto, errPuerto := appPuerto.leer()
	baseDatos, errBaseDatos := cargarBaseDatos()
	zona, errZona := cargarZonaHoraria()

	c := Config{
		Entorno:     entorno,
		PuertoHTTP:  puerto,
		BaseDatos:   baseDatos,
		ZonaHoraria: zona,
	}

	return c, errors.Join(errEntorno, errPuerto, errBaseDatos, errZona)
}
