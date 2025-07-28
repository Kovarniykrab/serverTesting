package handlers

// SuccessResponse - успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse - ошибка
type ErrorResponse struct {
	Error string `json:"error"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	Message string `json:"message"`
}
