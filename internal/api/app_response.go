package api

type AppResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
}

func NewAppResponse(status int, message string, data any) *AppResponse {
	return &AppResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Success: true,
	}
}
