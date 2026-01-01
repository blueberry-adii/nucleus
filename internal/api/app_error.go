package api

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Errors  []error `json:"errors"`
	Success bool    `json:"success"`
}

func NewAppError(w http.ResponseWriter, status int, message string, errors []error) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&AppError{
		Status:  status,
		Message: message,
		Errors:  errors,
		Success: false,
	})
}
