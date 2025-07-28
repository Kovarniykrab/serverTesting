package handlers

import (
	"encoding/json"

	"github.com/Kovarniykrab/serverTesting/myApp/domain"
	"github.com/valyala/fasthttp"
)

// RegisterUser godoc
// @Summary     регистрация
// @Description регистрация нового пользователя
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.RegisterUserForm  true  "данные для регистрации"
// @Success      201  {object} SuccessResponse "Пользователь успешно зарегестрирован"
// @Failure      400  {object}  ErrorResponse "Неверный формат данных"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      409  {object}  ErrorResponse "Пользователь с таким Email/Login уже существует"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router      /api/user/register [POST].
func RegisterUserHandler(ctx *fasthttp.RequestCtx) {
	var user domain.RegisterUserForm

	if !ctx.IsPost() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("неверный формат данных")
		return
	}

	if user.Email == "" || user.Password == "" {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusConflict)
		ctx.WriteString("Email и пароль обязательны")
		return
	}

	// проверка занят ли login и email и телефон

	if user.ConfirmPassword != user.Password {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("пароли должны совпадать!")
		return
	}

	//генерация хеш из пароля

	//запись данных в бд

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.WriteString("Пользователь успешно зарегистрирован")
}

// DeleteUser godoc
// @Summary     Удаление пользователя
// @Description Удаление пользователя из системы
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.User  true  "Данные пользователя"
// @Success      200  {object} DeleteResponse "Пользователь успешно удален"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      409  {object}  ErrorResponse "Пользователь с таким ID не существует"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/delete{id} [DELETE].
func DeleteUserHandler(ctx *fasthttp.RequestCtx) {

	if !ctx.IsDelete() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}

	id := ctx.UserValue("id").(string)
	if id == "" {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("ID пользователя не указан")
		return
	}

	//проверка авторизации пользователя

	// проверка существования пользователя по id

	//подтверждение пароля перед удалением

	//удаление пользователя

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("Пользователь успешно удален")
}

// UpdatePassword godoc
// @Summary     изменение пользователя
// @Description изменение пароля пользователя
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.User  true  "Данные пользователя"
// @Success      200  {object} UpdateResponse "Пользователь успешно изменен"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      409  {object}  ErrorResponse "Пользователь с таким ID не существует"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/update [PUT].
func UpdatePasswordHandler(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPut() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}

	id := ctx.UserValue("id").(string)
	if id == "" {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("ID пользователя не указан")
		return
	}

	//проверка существования пользователя по id

	//обновление данных
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("Пользователь успешно изменен")

}

// AuthUSer godoc
// @Summary    Авторизация
// @Description Авторизация пользователя
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.UserAuthForm  true  "Данные для авторизации пользователя"
// @Success      200  {object} AuthResponse "Успешный вход"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      409  {object}  ErrorResponse "Неверный Email/Password"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/login [POST].
func AuthUserHandler(ctx *fasthttp.RequestCtx) {
	var user domain.UserAuthForm

	if !ctx.IsPost() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("Неверный формат данных")
		return
	}

	//проверка существования польхователя в бд по логену или email

	//проверка пароля

	//генерация токена

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	//выдать сгенерированный jwt токен
	//сообщение "Успешный вход"

}

// GetUser godoc
// @Summary     Поиск польхователя
// @Description Поиск пользователя по ID
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.User  true  "Данные пользователя"
// @Success      200  {object} GetUserResponse "Пользователь найден"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/profile/{id} [GET].
func GetUserHandler(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}
	id := ctx.UserValue("id").(string)
	if id == "" {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("ID пользователя не указан")
		return
	}

	//проверка существования пользователя в бд

	//выдать результат

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("Пользователь найден")

}

// LogoutUser godoc
// @Summary     выход
// @Description окончание сессии
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param object  body  domain.User  true  "Данные пользователя"
// @Failure      400  {object}  ErrorResponse "Неверный запрос"
// @Failure      405  {object}  ErrorResponse "Метод не разрешен"
// @Failure      500  {object}  ErrorResponse "Ошибка сервера"
// @Router     /api/user/logout [POST].
func LogoutUserHandler(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.WriteString("метод не разрешён")
		return
	}

	// проверка сущестования токена. Если токен просрочен - выход

	//удаление токена

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("Успешный выход")
}
