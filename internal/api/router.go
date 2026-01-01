package api

import "net/http"

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

func (r *Router) Group(pattern string) *Router {
	return &Router{
		group: r.group + pattern,
		mux:   r.mux,
	}
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	r.mux.Handle(r.group+pattern, Logging(handler))
}

func HealthRoutes(mux *http.ServeMux) {
	router := NewRouter(mux).Group("/api/v1/health")

	router.Handle("/", Health)
}
