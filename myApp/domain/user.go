package domain

type User struct {
	ID          int    `bun:"id" json:"id"`
	Email       string `bun:"email" json:"email"`
	Login       string `bun:"login" json:"login"`
	Password    string `bun:"password" json:"password"`
	PhoneNumber string `bun:"phone_number" json:"phone_number"`
}

// форма регистрации
type RegisterUserForm struct {
	// Почта
	Email string `json:"email"`
	// логин
	Login string `json:"login"`
	//пароль
	Password string `json:"password"`
	//подтверждение пароля
	ConfirmPassword string `json:"confirm_password"`
	//номер телефона
	PhoneNumber string `json:"phone_number"`
}

// форма авторизации
type UserAuthForm struct {
	// логин
	Login string `json:"login"`
	//пароль
	Password string `json:"password"`
}

// форма смены пароля
type ChangePassForm struct {
	//cтарый пароль
	OldPassword string `json:"old_password"`
	//новый пароль
	Password string `json:"password"`
	//подтверждение пароля
	ConfirmPass string `json:"confirm_pass"`
}

func (r *RegisterUserForm) ToUser() *User {
	return &User{
		Email:       r.Email,
		Login:       r.Login,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
	}
}
