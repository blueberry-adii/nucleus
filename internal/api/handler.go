package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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
	_, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	NewAppResponse(w, http.StatusOK, "API Healthy and Working", nil)
}

// Auth Endpoint
// -------------------------------------------------------------------------------------

type UserHandler struct {
	service auth.UserServiceCreator
}

type UserRequestBody struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(service auth.UserServiceCreator) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body UserRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		NewAppError(w, http.StatusBadRequest, "Invalid Request Body", []error{err})
		return
	}
	if err := h.service.CreateUser(ctx, body.Email, body.Name, body.Password); err != nil {
		var status int
		if auth.HandleDBError(err).Number == 1062 {
			status = http.StatusConflict
		} else {
			status = http.StatusInternalServerError
		}
		NewAppError(w, status, "Failed to signup user: "+err.Error(), []error{err})
		return
	}

	NewAppResponse(w, 200, "Signed up user successfully", nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body LoginRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		NewAppError(w, http.StatusBadRequest, "Invalid Request Body", []error{err})
		return
	}

	res, err := h.service.AuthenticateUser(ctx, body.Email, body.Password)
	if err != nil {
		NewAppError(w, http.StatusUnauthorized, "Invalid Email or Password", err)
		return
	}

	NewAppResponse(w, 200, "Logged In successfully", res)
}

// -------------------------------------------------------------------------------------
