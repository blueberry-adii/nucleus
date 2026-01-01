package api

import (
	"database/sql"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/auth"
)

// Custom Router for grouping patterns
type Router struct {
	handler Handler
	group   string
	mux     *http.ServeMux
}

func NewRouter(mux *http.ServeMux, handler Handler) *Router {
	return &Router{
		group:   "",
		mux:     mux,
		handler: handler,
	}
}

func (r *Router) Group(pattern string) *Router {
	if pattern == "/" {
		pattern = ""
	}
	return &Router{
		group: r.group + pattern,
		mux:   r.mux,
	}
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	r.mux.Handle(r.group+pattern, handler)
}

func HealthRoutes(mux *http.ServeMux) {
	handler := NewHealthHandler()
	router := NewRouter(mux, handler).Group("/api/v1/health")

	router.Handle("/", handler.Health)
}

func AuthRoutes(mux *http.ServeMux, db *sql.DB) {
	store := auth.NewMySqlStore(db)
	service := auth.NewUserService(store)
	handler := NewUserHandler(service)
	router := NewRouter(mux, handler).Group("/api/v1/auth")

	router.Handle("/signup", handler.Signup)
	router.Handle("/login", handler.Login)
}
