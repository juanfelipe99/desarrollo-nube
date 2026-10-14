package config

import (
	"fmt"
	"os"
	"strings"
)

type variable struct {
	nombre      string
	porDefecto  string
	obligatoria bool
}

func obligatoria(nombre string) variable {
	return variable{nombre: nombre, obligatoria: true}
}

func opcional(nombre, porDefecto string) variable {
	return variable{nombre: nombre, porDefecto: porDefecto}
}

// leer devuelve el valor sin espacios alrededor.
// Si está vacía, da error en las obligatorias y porDefecto en las opcionales.
func (v variable) leer() (string, error) {
	if valor := strings.TrimSpace(os.Getenv(v.nombre)); valor != "" {
		return valor, nil
	}

	if v.obligatoria {
		return "", fmt.Errorf("%s es obligatoria", v.nombre)
	}

	return v.porDefecto, nil
}
