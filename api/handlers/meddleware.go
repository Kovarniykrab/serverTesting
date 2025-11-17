package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func (app *App) AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		token := string(ctx.Request.Header.Cookie("session_token"))
		if token == "" {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, errors.New("session token is required"))
			return
		}

		claims, err := app.Service.JWTService.ValidateJwt(token)
		if err != nil {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, err)
			return
		}

		user, err := strconv.Atoi(claims.Subject)
		if err != nil {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, errors.New("invalid user ID in token"))
			return
		}
		_, err = app.Service.GetById(ctx, user)
		if err != nil {
			app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, errors.New("user not found"))
			return
		}

		ctx.SetUserValue(userCTX, user)

		next(ctx)
	}
}

func (app *App) PanicRecoveryMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				app.logs.Error("PANIC recovered", "error", err, "url", ctx.Request.URI(), "method", ctx.Method())
				app.sendErrorResponse(ctx, fasthttp.StatusInternalServerError, errors.New("internal server error"))
			}
		}()
		next(ctx)
	}
}

func (app *App) LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		app.logs.Info("Request started",
			"method", string(ctx.Method()),
			"url", string(ctx.Request.URI().FullURI()),
			"remote_addr", ctx.RemoteAddr(),
		)

		next(ctx)

		app.logs.Info("Request completed",
			"method", string(ctx.Method()),
			"url", string(ctx.Request.URI().FullURI()),
			"status", ctx.Response.StatusCode(),
			"duration", time.Since(start),
		)
	}
}

func (app *App) CorsMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", app.cfg.Cors)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}

		next(ctx)
	}
}

//validateJwt справить
//jwt_service убрать в отдельную папку
// ложить пользователя в контекст в мидлвар
