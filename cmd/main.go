package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go-mux-template/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
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
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
