package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"
)

func (app *App) AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		token := string(ctx.Request.Header.Cookie("session_token"))
		if token == "" {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, errors.New("session token is required"))
			return
		}

		userID, err := app.Service.JWTService.ValidateJwt(token)
		if err != nil {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, err)
			return
		}

		ctx.SetUserValue("userID", userID)

		next(ctx)
	}
}
