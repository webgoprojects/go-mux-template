package main

import (
	"net/http"
	"os"
	"path/filepath"

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

	// Serve static files
	staticDir := filepath.Join(baseDir, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Start the server
	logger.Logger.Info("Server starting",
		zap.String("port", cfg.Port),
		zap.String("environment", cfg.Environment),
	)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
