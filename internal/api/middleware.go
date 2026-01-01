package api

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		log.Printf("%s %s %v", r.Method, r.URL.Path, elapsed)
	}
}
