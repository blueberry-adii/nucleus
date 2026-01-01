package main

import (
	"log"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/api"
)

func main() {
	api.HealthRoutes()

	log.Printf("Server running on port :8080")
	http.ListenAndServe(":8080", api.Mux)
}
