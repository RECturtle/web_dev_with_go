package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(filepath)

	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tPath)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port: 3000")
	http.ListenAndServe(":3000", r)
}
