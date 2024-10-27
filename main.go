package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler responds with "Hello, World!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// healthHandler responds with a simple 200 OK status to indicate the service is healthy
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status: Healthy")
}

// liveHandler responds with a simple 200 OK status to indicate the service is live
func liveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Status: Live")
}

func main() {
	http.HandleFunc("/hello", helloHandler)   // Route for /hello
	http.HandleFunc("/health", healthHandler) // Route for health check
	http.HandleFunc("/live", liveHandler)     // Route for liveness check

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
