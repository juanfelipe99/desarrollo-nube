package config

import (
	"strings"
	"testing"
)

// partesPostgres define todas las partes de la conexión, sin DATABASE_URL.
func partesPostgres(t *testing.T) {
	t.Helper()

	t.Setenv(databaseURL.nombre, "")
	t.Setenv(postgresUser.nombre, "app")
	t.Setenv(postgresPassword.nombre, "clave")
	t.Setenv(postgresHost.nombre, "localhost")
	t.Setenv(postgresPuerto.nombre, "")
	t.Setenv(postgresDB.nombre, "seguimiento")
	t.Setenv(postgresSSLMode.nombre, "")
}

func TestCargarBaseDatosArmaLaURL(t *testing.T) {
	partesPostgres(t)

	url, err := cargarBaseDatos()

	esperada := "postgres://app:clave@localhost:" + postgresPuerto.porDefecto + "/seguimiento?sslmode=" + postgresSSLMode.porDefecto
	if err != nil || url != esperada {
		t.Errorf("cargarBaseDatos() = %q, %v, se esperaba %q", url, err, esperada)
	}
}

func TestCargarBaseDatosPrefiereDatabaseURL(t *testing.T) {
	partesPostgres(t)
	t.Setenv(databaseURL.nombre, "postgres://otro@nube:5432/bd")

	url, err := cargarBaseDatos()

	if err != nil || url != "postgres://otro@nube:5432/bd" {
		t.Errorf("cargarBaseDatos() = %q, %v, debería usar %s", url, err, databaseURL.nombre)
	}
}

func TestCargarBaseDatosExigeLasPartes(t *testing.T) {
	partesPostgres(t)
	t.Setenv(postgresUser.nombre, "")
	t.Setenv(postgresHost.nombre, "")

	_, err := cargarBaseDatos()

	if err == nil || !strings.Contains(err.Error(), postgresUser.nombre) || !strings.Contains(err.Error(), postgresHost.nombre) {
		t.Errorf("se esperaba un error que mencionara %s y %s: %v", postgresUser.nombre, postgresHost.nombre, err)
	}
}
