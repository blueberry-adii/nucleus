package main

import (
	"log"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/api"
)

func main() {

	mux := http.NewServeMux()
	handler := api.Logging(mux.ServeHTTP)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	api.HealthRoutes(mux)

	log.Printf("Server running on port :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
