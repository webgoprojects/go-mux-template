package main

import (
	"net/http"
	"os"
	"path/filepath"

	"go-mux-template/pkg/handlers"
	"go-mux-template/pkg/logger"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	// Initialize structured logger
	if err := logger.Init("info"); err != nil {
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
	logger.Logger.Info("Server starting", zap.String("port", "8080"))
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
