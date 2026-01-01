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
	"github.com/blueberry-adii/nucleus.git/internal/platform/database"
	"github.com/blueberry-adii/nucleus.git/internal/platform/shutdown"
)

func main() {
	mux := http.NewServeMux()

	config := database.Config{
		User:     "root",
		Password: "pass",
		Host:     "localhost",
		Port:     3306,
		Database: "nucleus",
	}
	db, err := database.NewMySQL(config)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	api.InitRoutes(mux)

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
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	shutdown.Run(srv)
}
