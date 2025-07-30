package domain

type User struct {
	ID          int    `bun:"id" json:"id"`
	Email       string `bun:"email" json:"email"`
	Name        string `bun:"name" json:"name"`
	Password    string `bun:"password" json:"password"`
	PhoneNumber string `bun:"phone_number" json:"phone_number"`
}

// форма регистрации
// аннотация на обязательные поля!!!
type RegisterUserForm struct {
	// Почта
	Email string `json:"email"`
	//пароль
	Name     string `json:"name"`
	Password string `json:"password"`
	//подтверждение пароля
	ConfirmPassword string `json:"confirm_password"`
	//номер телефона
	PhoneNumber string `json:"phone_number"`
}

// форма авторизации
type UserAuthForm struct {
	// логин
	Email string `json:"email"`
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
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
	}
}
