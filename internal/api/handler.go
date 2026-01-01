package api

import (
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/auth"
)

type Handler interface {
}

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Healthy and Working"))
}

type UserHandler struct {
	service auth.UserServiceCreator
}

func NewUserHandler(service auth.UserServiceCreator) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {

}
