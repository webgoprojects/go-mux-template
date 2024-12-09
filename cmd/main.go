package main

import (
        "net/http"

        "github.com/gorilla/mux"
)

func main() {
        r := mux.NewRouter()

        // Define routes and handlers
        r.HandleFunc("/", homeHandler)
        r.HandleFunc("/about", aboutHandler)

        // Serve static files
        r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

        // Start the server
        http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
        // Render the index.html template
        // ...
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
        // Render the about.html template
        // ...
}