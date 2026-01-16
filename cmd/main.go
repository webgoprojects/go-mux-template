package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"go-mux-template/pkg/config"
	"go-mux-template/pkg/handlers"
	"go-mux-template/pkg/logger"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	// Load configuration from environment variables
	cfg := config.Load()

	// Initialize structured logger with config
	if err := logger.Init(cfg.LogLevel); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	r := mux.NewRouter()

	// Get the base directory (assuming we run from project root)
	baseDir, err := os.Getwd()
	if err != nil {
		logger.Logger.Fatal("Failed to get working directory", zap.Error(err))
	}

	// Define routes and handlers
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	// Serve static files
	staticDir := filepath.Join(baseDir, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Create HTTP server with timeouts
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Logger.Info("Server starting",
			zap.String("port", cfg.Port),
			zap.String("environment", cfg.Environment),
			zap.String("log_level", cfg.LogLevel),
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Logger.Info("Server shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Logger.Info("Server exited")
}
