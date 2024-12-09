package handlers

import (
        "net/http"
        "text/template"
)

var templates = template.Must(template.ParseFiles("templates/base.html", "templates/index.html", "templates/about.html"))

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