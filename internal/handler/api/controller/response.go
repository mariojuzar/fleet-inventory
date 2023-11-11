package controller

type JsonResponse[T any] struct {
	Data    T              `json:"data,omitempty"`
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Success bool           `json:"success,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
