package main

import (
	"fmt"
	"net/http"

	BA "BA/internal/Functions"
)

func main() {
	// Link static files such as CSS with the HTML
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if style.css is existed
		err := BA.EnsureFile(
			"/internal/frontend/style.css",
			"https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/frontend/style.CSS",
		)
		if err != nil {
			err := BA.EnsureFile("/internal/frontend/500.html", "https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/frontend/500.html")
			if err != nil {
				http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, "/internal/frontend/500.html")
			return
		}

		// Serve CSS to client
		http.FileServer(http.Dir("/internal/frontend")).ServeHTTP(w, r)
	})))

	// Handle wrong routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			err := BA.EnsureFile("/internal/frontend/404.html", "https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/frontend/404.html")
			if err != nil {
				http.Error(w, "404 Server Error", http.StatusNotFound)
				return
			}
			http.ServeFile(w, r, "/internal/frontend/404.html")
			return
		}

		BA.FormHandler(w, r)
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
