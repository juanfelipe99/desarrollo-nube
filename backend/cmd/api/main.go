package main

import (
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("servicio", "api")
	if err := ejecutar(log); err != nil {
		log.Error("la API terminó con error", "error", err)
		os.Exit(1)
	}
}

func ejecutar(log *slog.Logger) error {
	log.Info("API iniciada")
	return nil
}
