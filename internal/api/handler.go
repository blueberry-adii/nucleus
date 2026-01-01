package api

import (
	"encoding/json"
	"net/http"

	"github.com/blueberry-adii/nucleus.git/internal/auth"
)

type Handler interface {
}

// Health Endpoint
// -------------------------------------------------------------------------------------

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	res := NewAppResponse(http.StatusOK, "API Healthy and Working", nil)

	json.NewEncoder(w).Encode(res)
}

// Auth Endpoint
// -------------------------------------------------------------------------------------

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

// -------------------------------------------------------------------------------------
