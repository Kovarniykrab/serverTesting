package domain

type User struct {
	ID          int    `bun:"id" json:"id"`
	Email       string `bun:"email" json:"email"`
	Name        string `bun:"name" json:"name"`
	DateOfBirth string `bun:"date_of_birth" json:"date_of_birth"`
	Password    string `bun:"password" json:"password"`
}

// форма регистрации
// аннотация на обязательные поля!!!
type RegisterUserForm struct {
	// Почта
	// Поле обязательно
	Email string `json:"email"`
	//имя пользователя
	// Поле обязательно
	Name string `json:"name"`
	//дата рождения
	// Поле обязательно
	DateOfBirth string `json:"date_of_birth"`
	//пароль
	// Поле обязательно
	Password string `json:"password"`
	//подтверждение пароля
	// Поле обязательно
	ConfirmPassword string `json:"confirm_password"`
}

// форма авторизации
type UserAuthForm struct {
	// логин
	// Поле обязательно
	Email string `json:"email"`
	//пароль
	// Поле обязательно
	Password string `json:"password"`
}

// форма смены данных пользователя
type ChangeUserForm struct {
	//имя пользователя
	// Поле обязательно
	Name string `bun:"name" json:"name"`
	//дата рождения
	// Поле обязательно
	DateOfBirth string `bun:"date_of_birth" json:"date_of_birth"`
}

// форма смены пароля
type ChangePassForm struct {
	//cтарый пароль
	// Поле обязательно
	OldPassword string `json:"old_password"`
	//новый пароль
	// Поле обязательно
	Password string `json:"password" bun:"password"`
	//подтверждение пароля
	// Поле обязательно
	ConfirmPass string `json:"confirm_pass"`
}

type Users []User
