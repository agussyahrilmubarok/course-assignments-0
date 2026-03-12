package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Info().Msg("server started at :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	<-ctx.Done()

	log.Info().Msg("shutting down server")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()

	logger := log.With().
		Str("request_id", requestID).
		Str("path", r.URL.Path).
		Logger()

	ctx := logger.WithContext(r.Context())

	log.Ctx(ctx).Info().Msg("request received")

	name := r.URL.Query().Get("name")
	res, err := greeting(ctx, name)
	if err != nil {
		log.Ctx(ctx).
			Warn().
			Err(err).
			Msg("request failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(res))
}

func greeting(ctx context.Context, name string) (string, error) {
	log := log.Ctx(ctx)

	log.Info().
		Str("name", name).
		Msg("running greeting")

	if name == "" {
		err := fmt.Errorf("name is required")
		log.Warn().
			Err(err).
			Str("field", "name").
			Msg("validation failed")
		return "", err
	}

	if len(name) < 3 {
		err := fmt.Errorf("name must be at least 3 characters")
		log.Warn().
			Err(err).
			Str("field", "name").
			Int("min_length", 3).
			Int("actual_length", len(name)).
			Msg("validation failed")
		return "", err
	}

	return fmt.Sprintf("Hi %s", name), nil
}
