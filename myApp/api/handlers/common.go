package handlers

// SuccessResponse - успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse - ошибка
type ErrorResponse struct {
	Error string `json:"error"`
}

