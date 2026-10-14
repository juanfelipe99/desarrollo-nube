package config

import (
	"errors"
	"net"
	"net/url"
)

var (
	databaseURL      = opcional("DATABASE_URL", "")
	postgresUser     = obligatoria("POSTGRES_USER")
	postgresPassword = obligatoria("POSTGRES_PASSWORD")
	postgresHost     = obligatoria("POSTGRES_HOST")
	postgresPuerto   = opcional("POSTGRES_PUERTO", "5432")
	postgresDB       = obligatoria("POSTGRES_DB")
	postgresSSLMode  = opcional("POSTGRES_SSLMODE", "disable")
)

// cargarBaseDatos usa DATABASE_URL si está definida o la arma con las variables de postgres.
func cargarBaseDatos() (string, error) {
	if urlCompleta, _ := databaseURL.leer(); urlCompleta != "" {
		return urlCompleta, nil
	}

	usuario, errUsuario := postgresUser.leer()
	clave, errClave := postgresPassword.leer()
	host, errHost := postgresHost.leer()
	puerto, errPuerto := postgresPuerto.leer()
	bd, errBD := postgresDB.leer()
	sslMode, errSSLMode := postgresSSLMode.leer()

	if err := errors.Join(errUsuario, errClave, errHost, errPuerto, errBD, errSSLMode); err != nil {
		return "", err
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(usuario, clave),
		Host:     net.JoinHostPort(host, puerto),
		Path:     "/" + bd,
		RawQuery: url.Values{"sslmode": {sslMode}}.Encode(),
	}

	return u.String(), nil
}
