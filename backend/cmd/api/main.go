package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata" // zonas horarias dentro del binario, la imagen de Docker no las trae

	"github.com/gin-gonic/gin"

	"seguimiento/internal/config"
	"seguimiento/internal/handler"
)

const tiempoMaxApagado = 15 * time.Second

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

	if cfg.EsProduccion() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Se cancela con Ctrl+C o cuando Docker detiene el contenedor.
	ctx, detener := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer detener()

	servidor := &http.Server{
		Addr:              ":" + cfg.PuertoHTTP,
		Handler:           handler.NuevoRouter(handler.Dependencias{}),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		IdleTimeout:       2 * time.Minute,
	}

	errServidor := make(chan error, 1)

	go func() {
		log.Info("API escuchando", "puerto", cfg.PuertoHTTP, "entorno", cfg.Entorno)
		errServidor <- servidor.ListenAndServe()
	}()

	select {
	case err := <-errServidor:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

	case <-ctx.Done():
		log.Info("deteniendo la API")
	}

	// Deja de aceptar peticiones y espera las que están en curso.
	ctxApagado, cancelar := context.WithTimeout(context.Background(), tiempoMaxApagado)
	defer cancelar()

	return servidor.Shutdown(ctxApagado)
}
