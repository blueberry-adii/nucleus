package api

import "net/http"

var Mux = http.NewServeMux()

// Custom Router for grouping patterns
type Router struct {
	group string
}

func NewRouter() *Router {
	return &Router{
		group: "",
	}
}

func (r *Router) Group(pattern string) *Router {
	r.group = pattern
	return r
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	Mux.Handle(r.group+pattern, Logging(handler))
}

var router = NewRouter()

func HealthRoutes() {
	router.Group("/api/v1/health")

	router.Handle("/", Health)
}
