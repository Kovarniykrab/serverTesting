package handlers

import (
	"encoding/json"

	"github.com/Kovarniykrab/serverTesting/myApp/domain"
	"github.com/valyala/fasthttp"
)

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
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("Email и пароль обязательны")
		return
	}

	// проверка занят ли login и email

	if user.Password == user.ConfirmPassword {

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

}

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
