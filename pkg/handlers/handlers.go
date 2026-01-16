package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"go-mux-template/pkg/logger"

	"go.uber.org/zap"
)

var templates *template.Template

func init() {
	// Get the base directory (assuming we run from project root)
	baseDir, err := os.Getwd()
	if err != nil {
		panic("Failed to get working directory: " + err.Error())
	}

	templateDir := filepath.Join(baseDir, "templates")
	templates = template.Must(template.ParseFiles(
		filepath.Join(templateDir, "base.html"),
		filepath.Join(templateDir, "index.html"),
		filepath.Join(templateDir, "about.html"),
	))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("Home handler called",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
	)

	if err := templates.ExecuteTemplate(w, "base", map[string]interface{}{
		"Title":   "Home Page",
		"Content": "This is the home page.",
	}); err != nil {
		logger.Logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("About handler called",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
	)

	if err := templates.ExecuteTemplate(w, "base", map[string]interface{}{
		"Title":   "About Page",
		"Content": "This is the about page.",
	}); err != nil {
		logger.Logger.Error("Failed to execute template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
