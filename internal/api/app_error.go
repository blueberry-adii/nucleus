package api

type AppError struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Errors  []error `json:"errors"`
	Success bool    `json:"success"`
}

func NewAppError(status int, message string, errors []error) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
		Errors:  errors,
		Success: false,
	}
}
