package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to My Simple HTTP Server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This server was handcrafter with Go!")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Searching for: %s", query)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/helo/")
	if path == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s", path)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", start.Format(time.RFC3339), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Request processed in %s\n", time.Since(start))
	})
}

func main() {
	// Set up handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/helo/", helloHandler)

	// wrap handlers with middleware
	loggedMux := loggingMiddleware(mux)

	// Start server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
