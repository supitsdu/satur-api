package server

import (
	"log"
	"net/http"
)

// Listen starts the HTTP server on the specified address with the given handler
func Listen(address string, handler http.Handler) {
	log.Printf("Server listening on %s\n", address)

	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
