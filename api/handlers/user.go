package handlers

import (
	"encoding/json"
	"log/slog"
	"strconv"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
	"github.com/valyala/fasthttp"
)

// RegisterUser godoc
// @Summary     регистрация
// @Description во время регистации хендлер принимает данные, которые подает пользователь и проверяет
// @Description если email не занят,
// @Description то данные записываются на сервер в базу данных, пользователю присваивается id
// @Description и отправляется письмо на почту, с просьбой подтвердить аккаунт и ссылкой, по которой нужно перейти для подтверждения аккаунта
// @Description если занят email, на указанную почту приходит уведомление о том, что на его почту пытаются зарегестрировать новый аккаунт
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.RegisterUserForm  true  "Обязательные поля : email, password, name, confirm_password, date_of_birth"
// @Success      204  "Пользователь успешно зарегестрирован"
// @Failure      400  {object}  ErrorResponse "Неверный формат данных"
// @Failure      404  {object}  ErrorResponse "Неверный запрос"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router      /api/user/register [POST]
func (app *App) RegisterUserHandler(ctx *fasthttp.RequestCtx) {
	var form domain.RegisterUserForm

	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	if err := app.Service.RegisterUser(ctx, form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, err)
		return
	}

	app.sendSuccessResponse(ctx, fasthttp.StatusCreated, "Пользователь успешно зарегистрирован")

	//запись данных в бд
}

// DeleteUser godoc
// @Summary     Удаление пользователя
// @Description Пользователя удаляют из системы по ID. Хендлер принимает ID,
// @Description и с его помощью находит пользователя в базе данных и удаляет его из нее
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Обязательные поля: id"
// @Success      204  "Пользователь успешно удален"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      404  {object}  ErrorResponse "запрашиваемая страница не существует"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/delete/{id} [DELETE]
func (app *App) DeleteUserHandler(ctx *fasthttp.RequestCtx) {

	idStr := ctx.UserValue("id").(string)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.NotFound(err))
		return
	}

	if err := app.Service.DeleteUser(ctx, id); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusNotFound, err)
		return
	}

	app.sendSuccessResponse(ctx, fasthttp.StatusOK, "Пользователь успешно удален")
}

// UpdatePassword godoc
// @Summary     изменение пользователя
// @Description хендлер проверяет авторизован ли пользователь и запрашивает подтверждение пароля
// @Description если все в порядке, пользователю подается форма, для изменения данных.
// @Description затем пользователь подает форму на сервер, и они записываются вместо старых
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.ChangePassForm  true  "Обязательные поля : password, confirm_password"
// @Success      204  "Пароль успешно изменен"
// @Failure      401 {object} ErrorResponse "Не авторизован"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/change_password/{id} [PUT]
func (app *App) UpdatePasswordHandler(ctx *fasthttp.RequestCtx) {

	idStr := ctx.UserValue("id").(string)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	var form domain.ChangePassForm
	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	if err := app.Service.UpdatePassword(ctx, id, form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, err)
		return
	}

	app.sendSuccessResponse(ctx, fasthttp.StatusOK, "Пароль успешно изменен")

}

// AuthUSer godoc
// @Summary    Авторизация
// @Description Авторизация происходит с помощью email и пароля
// @Description хендлер принимает почту и пароль. По почте ищет пользователя и сверяет 2 хеша. Если они совпадают - пользователь авторизуется
// @Description при авторизации пользоватcz проверяет временный jwt токен.
// @Description если авторизация успешна, вызываются хендлеры, которые в свою очередь выдают пользователю данные профиля и историю его диалогов, внутри которых
// @Description переписка с конкретным пользователем
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserAuthForm  true  "Обязательные поля : email, password"
// @Success      200  {object}  AuthResponse  "Успешный вход"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      404  {object}  ErrorResponse "Неверный Email/Password"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/login [POST]
func (app *App) AuthUserHandler(ctx *fasthttp.RequestCtx) {
	var user domain.UserAuthForm

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	userRender, err := app.Service.AuthUser(ctx, user)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, domain.Unauthorized(err))
		return
	}

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("session_token")
	cookie.SetValue(userRender.Token)
	cookie.SetExpire(time.Now().Add(24 * time.Hour))
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")
	ctx.Response.Header.SetCookie(cookie)

	//проверка существования польхователя в бд по логену или email

	//проверка пароля

	//генерация токена

	app.sendSuccessResponse(ctx, fasthttp.StatusOK, "Успешная авторизация")

	//выдать сгенерированный jwt токен
	//сообщение "Успешный вход"

}

// GetUser godoc
// @Summary     Поиск польхователя
// @Description хендлер получает ID или name и ищет пользователей по id или name в базе данных и выводит все результаты в форме списка.
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "ID пользователя"
// @Success      200  {object} SuccessResponse "Пользователь найден"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/profile/{id} [GET]
func (app *App) GetUserHandler(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	user, err := app.Service.GetUserById(ctx, id)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusNotFound, domain.NotFound(err))
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	jsonData, err := json.Marshal(user)
	if err != nil {
		slog.Error("Failed to marshal user response", "error", err)
		app.sendErrorResponse(ctx, fasthttp.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	if _, err := ctx.Write(jsonData); err != nil {
		slog.Error("Failed to write user response", "error", err)
	}
}

// ChangeUser godoc
// @Summary     изменение данных пользователя
// @Description хендлер принимает форму, в которой содержатся новые данные и данные оставшиеся без изменений.
// @Description он записывает все данные вместо старых
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.ChangeUserForm  true  "Данные пользователя"
// @Success      204  "Данные успешно обновлены"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/change_user/{id} [PUT]
func (app *App) ChangeUserHandler(ctx *fasthttp.RequestCtx) {

	idStr := ctx.UserValue("id").(string)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	var form domain.ChangeUserForm

	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, domain.BadRequest(err))
		return
	}

	if err := app.Service.UpdateUser(ctx, id, form); err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusBadRequest, err)
		return
	}

	app.sendSuccessResponse(ctx, fasthttp.StatusOK, "Данные успешно обновлены")
}

// LogoutUser godoc
// @Summary     выход
// @Description сессия завершается по сигналу или по истечении токена
// @Description пользователя должно выкинуть на страницу авторизации
// @Tags         USER
// @Accept       json
// @Produce      json
// @Success      204  "успешный выход"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/logout [POST]
func (app *App) LogoutUserHandler(ctx *fasthttp.RequestCtx) {

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("session_token")
	cookie.SetValue("")
	cookie.SetExpire(time.Now().Add(-1 * time.Hour))
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")

	ctx.Response.Header.SetCookie(cookie)
	app.sendSuccessResponse(ctx, fasthttp.StatusOK, "Успешный выход")
}

// Check godoc
// @Summary Проверка авторизации
// @Description Проверяет валидность токена и возвращает данные пользователя
// @Tags AUTH
// @Produce json
// @Success 204
// @Failure 401 {object} ErrorResponse "ошибка"
// @Router /api/user/check [GET]
func (app *App) CheckHandler(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue("userID").(int)

	userRender, err := app.Service.CheckUser(ctx, userID)
	if err != nil {
		app.sendErrorResponse(ctx, fasthttp.StatusUnauthorized, domain.Unauthorized(err))
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	jsonData, err := json.Marshal(userRender)
	if err != nil {
		slog.Error("Failed to marshal user response", "error", err)
		app.sendErrorResponse(ctx, fasthttp.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	if _, err := ctx.Write(jsonData); err != nil {
		slog.Error("Failed to write user response", "error", err)
	}
}
