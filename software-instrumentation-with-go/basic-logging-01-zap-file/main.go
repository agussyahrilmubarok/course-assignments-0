package main

import (
	"context"
	"errors"
	"example/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	// Initialize the optimized Zap logger
	logger.InitLogger()

	// Flush any buffered log entries before the application exits
	defer logger.Log.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		logger.Log.Info("server started", zap.String("port", ":8080"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("failed to start server", zap.Error(err))
		}
	}()

	// Wait for termination signal
	<-ctx.Done()

	logger.Log.Info("shutting down server gracefully")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Log.Error("failed to shutdown server", zap.Error(err))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()

	// Create a child logger with request context
	// Zap uses strongly typed fields for better performance (less allocation)
	reqLogger := logger.Log.With(
		zap.String("request_id", requestID),
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),
	)

	ctx := logger.WithCtx(r.Context(), reqLogger)

	reqLogger.Info("request received")

	name := r.URL.Query().Get("name")
	res, err := greeting(ctx, name)
	if err != nil {
		reqLogger.Warn("request failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(res))
}

func greeting(ctx context.Context, name string) (string, error) {
	log := logger.FromCtx(ctx)

	log.Info("running greeting logic", zap.String("input_name", name))

	if name == "" {
		err := errors.New("name is required")
		log.Warn("validation failed",
			zap.Error(err),
			zap.String("field", "name"),
		)
		return "", err
	}

	if len(name) < 3 {
		err := errors.New("name is too short")
		log.Warn("validation failed",
			zap.Error(err),
			zap.Int("min_length", 3),
			zap.Int("actual_length", len(name)),
		)
		return "", err
	}

	return fmt.Sprintf("Hi %s", name), nil
}
