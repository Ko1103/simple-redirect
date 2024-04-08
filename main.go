package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
			// robots.txt that disallows all robots.
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprint(w, "User-agent: *\nDisallow: /")
		})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectURL := os.Getenv("REDIRECT_URL")
		if r.URL.Path != "/health" {
			http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
