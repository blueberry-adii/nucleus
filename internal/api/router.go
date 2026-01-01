package api

import (
	"database/sql"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/auth"
)

// Custom Router for grouping patterns
type Router struct {
	group string
	mux   *http.ServeMux
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{
		group: "",
		mux:   mux,
	}
}

func HealthRoutes(mux *http.ServeMux) {
	handler := NewHealthHandler()
	router := NewRouter(mux).Group("/api/v1/health")

	router.Get("/", handler.Health)
}

func AuthRoutes(mux *http.ServeMux, db *sql.DB) {
	store := auth.NewMySqlStore(db)
	service := auth.NewUserService(store)
	handler := NewUserHandler(service)
	router := NewRouter(mux).Group("/api/v1/auth")

	router.Post("/signup", handler.Signup)
	router.Post("/login", handler.Login)
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

func (r *Router) Handle(pattern string, handler http.HandlerFunc, method string) {
	if pattern == "/" {
		pattern = ""
	}
	r.mux.HandleFunc(r.group+pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler, "GET")
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler, "POST")
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler, "PUT")
}

func (r *Router) Patch(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler, "PATCH")
}

func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler, "DELETE")
}
