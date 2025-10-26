package handlers

import (
	"encoding/json"
	"log/slog"

	"github.com/valyala/fasthttp"
)

// SuccessResponse - успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse - ошибка
type ErrorResponse struct {
	Error string `json:"error"`
}

// AuthResponse - токен
type AuthResponse struct {
	JWTToken string `json:"JWT"`
}

func (app *App) sendErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, err error) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	response := ErrorResponse{Error: errorMsg}
	jsonData, err := json.Marshal(response)
	if err != nil {
		slog.Error("Failed to marshal error response", "error", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		if _, writeErr := ctx.Write([]byte(`{"error": "internal server error"}`)); writeErr != nil {
			slog.Error("Failed to write fallback error response", "error", writeErr)
		}
		return
	}

	if _, err := ctx.Write(jsonData); err != nil {
		slog.Error("Failed to write error response", "error", err)
	}
}

func (app *App) sendSuccessResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	response := SuccessResponse{Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		slog.Error("Failed to marshal success response", "error", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		if _, writeErr := ctx.Write([]byte(`{"error": "internal server error"}`)); writeErr != nil {
			slog.Error("Failed to write fallback error response", "error", writeErr)
		}
		return
	}

	if _, err := ctx.Write(jsonData); err != nil {
		slog.Error("Failed to write success response", "error", err)
	}
}
