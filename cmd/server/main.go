package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blueberry-adii/nucleus.git/internal/api"
)

func main() {

	mux := http.NewServeMux()

	api.HealthRoutes(mux)
	api.AuthRoutes(mux)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	handler := api.Logging(mux.ServeHTTP)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server running on port :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	<-ctx.Done()
}
