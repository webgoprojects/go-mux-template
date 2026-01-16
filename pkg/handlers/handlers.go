package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
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
        templates.ExecuteTemplate(w, "base", map[string]interface{}{
                "Title": "Home Page",
                "Content": "This is the home page.",
        })
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
        templates.ExecuteTemplate(w, "base", map[string]interface{}{
                "Title": "About Page",
                "Content": "This is the about page.",
        })
}
