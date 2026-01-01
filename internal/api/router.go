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

var initRouter *Router

// ROUTES
// --------------------------------------------------------------------------------------------------------
func InitRoutes(mux *http.ServeMux) {
	initRouter = NewRouter(mux).Group("/api/v1")
	HealthRoutes(mux)
}

func HealthRoutes(mux *http.ServeMux) {
	handler := NewHealthHandler()
	router := initRouter.Group("/health")

	router.Get("/", handler.Health)
}

func AuthRoutes(mux *http.ServeMux, db *sql.DB) {
	store := auth.NewMySqlStore(db)
	service := auth.NewUserService(store)
	handler := NewUserHandler(service)
	router := initRouter.Group("/auth")

	router.Post("/signup", handler.Signup)
	router.Post("/login", handler.Login)
}

// Used to group path patterns
// --------------------------------------------------------------------------------------------------------
func (r *Router) Group(pattern string) *Router {
	if pattern == "/" {
		pattern = ""
	}
	return &Router{
		group: r.group + pattern,
		mux:   r.mux,
	}
}

// Handles based on HTTP Method
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

// HTTP METHODS
// --------------------------------------------------------------------------------------------------------
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
