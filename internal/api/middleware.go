package api

import (
	"log"
	"net/http"
	"time"

	"github.com/blueberry-adii/nucleus.git/internal/auth"
	"github.com/golang-jwt/jwt/v5"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		log.Printf("%s %s %v", r.Method, r.URL.Path, elapsed)
	}
}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			NewAppError(w, http.StatusUnauthorized, "Unauthorized: no token cookie", []error{err})
			return
		}

		token, err := auth.VerifyJWT(cookie.Value)
		if err != nil {
			NewAppError(w, http.StatusUnauthorized, "Unauthorized: invalid token", []error{err})
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			NewAppError(w, http.StatusUnauthorized, "Unauthorized: invalid claims", nil)
			return
		}

		next.ServeHTTP(w, r)
	}
}
