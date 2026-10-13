package main

import (
	"log/slog"
	"os"

	_ "time/tzdata"

	"seguimiento/internal/config"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("servicio", "api")
	if err := ejecutar(log); err != nil {
		log.Error("la API terminó con error", "error", err)
		os.Exit(1)
	}
}

func ejecutar(log *slog.Logger) error {
	cfg, err := config.Cargar()
	if err != nil {
		return err
	}

	log.Info(
		"configuración cargada",
		"entorno", cfg.Entorno,
		"puerto", cfg.PuertoHTTP,
		"zona_horaria", cfg.ZonaHoraria.String(),
	)

	return nil
}
