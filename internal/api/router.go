package api

import "net/http"

var Mux = http.NewServeMux()

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
	r.group = pattern
	return r
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	r.mux.Handle(r.group+pattern, Logging(handler))
}

func HealthRoutes() {
	var router = NewRouter(Mux)

	router.Group("/api/v1/health")

	router.Handle("/", Health)
}
