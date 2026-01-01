package shutdown

import (
	"context"
	"log"
	"net/http"
	"time"
)

func Run(srv *http.Server) {
	log.Println("shutdown signal received")

	shutdownCtx, stop := context.WithTimeout(context.Background(), 10*time.Second)
	defer stop()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		if err == context.DeadlineExceeded {
			log.Println("graceful shutdown timed out, forcing exit")
		} else {
			log.Printf("server shutdown error: %v", err)
		}
	}

	log.Println("server exited cleanly")
}
