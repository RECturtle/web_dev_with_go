package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my amazing site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h1>Wecome to the contact page!</h1>
	<p>To get in touch email me at <a href="mailto:spencer.mertes@gmail.com">spencer.mertes@gmail.com</a>.
	`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h1>Frequently Asked Questions</h1>
	<p>
	<b>Q: Is there a free version?</b>
	<li>Yes! We offer a 30 day free trial</li>
	</p>
	<p>
	<b>Q: Is this the coolest website?</b>
	<li>Of course, no question about it</li>
	</p>
	`)
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	param := chi.URLParam(r, "param")
	fmt.Fprint(w, "<p>Welcome user: ", param, "</p>")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/param/{param}", paramHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port: 3000")
	http.ListenAndServe(":3000", r)
}
