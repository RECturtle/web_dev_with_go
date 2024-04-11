package main

import (
	"fmt"
	"net/http"
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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/path", pathHandler)

	fmt.Println("Starting the server on port: 3000")
	http.ListenAndServe(":3000", router)
}
