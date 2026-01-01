package api

import (
	"encoding/json"
	"net/http"
)

type AppResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
}

func NewAppResponse(w http.ResponseWriter, status int, message string, data any) {
	json.NewEncoder(w).Encode(&AppResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Success: true,
	})
}
