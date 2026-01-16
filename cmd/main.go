package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go-mux-template/pkg/config"
	"go-mux-template/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from environment variables
	cfg := config.Load()

	r := mux.NewRouter()

	// Get the base directory (assuming we run from project root)
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	// Define routes and handlers
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)

	// Serve static files
	staticDir := filepath.Join(baseDir, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Start the server
	log.Printf("Server starting on :%s (environment: %s)", cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
