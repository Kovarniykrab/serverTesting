package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// SuccessResponse - успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse - ошибка
type ErrorResponse struct {
	Error error `json:"error"`
}

// AuthResponse - токен
type AuthResponse struct {
	JWTToken string `json:"JWT"`
}

func (app *App) sendErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, err error) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	response := ErrorResponse{Error: err}
	if jsonData, err := json.Marshal(response); err == nil {
		ctx.Write(jsonData)
	}
}

func (app *App) sendSuccessResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	response := SuccessResponse{Message: message}
	if jsonData, err := json.Marshal(response); err == nil {
		ctx.Write(jsonData)
	}
}
