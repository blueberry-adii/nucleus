package main

import (
	"log"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/api"
)

func main() {

	mux := http.NewServeMux()

	api.HealthRoutes(mux)

	log.Printf("Server running on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
