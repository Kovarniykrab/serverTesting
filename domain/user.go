package domain

import (
	"strings"
	"time"

	"github.com/uptrace/bun"
)

const (
	MinPasswordLenght = 6
	MaxPasswordLenght = 40
	MinNameLength     = 6
	MaxNameLength     = 20
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID          int       `bun:"id,pk,autoincrement" json:"id"`
	Email       string    `bun:"email" json:"email"`
	Name        string    `bun:"name" json:"name"`
	DateOfBirth string    `bun:"date_of_birth" json:"date_of_birth"`
	Password    string    `bun:"password" json:"password"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
}

// форма регистрации
// аннотация на обязательные поля!!!
type RegisterUserForm struct {
	// Почта
	// Поле обязательно
	Email string `json:"email" binding:"required"`
	//имя пользователя
	// Поле обязательно
	Name string `json:"name" binding:"required"`
	//дата рождения
	// Поле обязательно
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	//пароль
	// Поле обязательно
	Password string `json:"password" binding:"required"`
	//подтверждение пароля
	// Поле обязательно
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// форма авторизации
type UserAuthForm struct {
	// логин
	// Поле обязательно
	Email string `json:"email" binding:"required"`
	//пароль
	// Поле обязательно
	Password string `json:"password" binding:"required"`
}

type UserRender struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

// форма смены данных пользователя
type ChangeUserForm struct {
	//имя пользователя
	// Поле обязательно
	Name string `bun:"name" json:"name" binding:"required"`
	//дата рождения
	// Поле обязательно
	DateOfBirth string `bun:"date_of_birth" json:"date_of_birth" binding:"required"`
}

// форма смены пароля
type ChangePassForm struct {
	//cтарый пароль
	// Поле обязательно
	OldPassword string `json:"old_password" binding:"required"`
	//новый пароль
	// Поле обязательно
	Password string `json:"password" bun:"password" binding:"required"`
	//подтверждение пароля
	// Поле обязательно
	ConfirmPass string `json:"confirm_pass" binding:"required"`
}

type Users []User

func isValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	if email == "" {
		return false
	}

	// Простая проверка email
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if parts[0] == "" || parts[1] == "" {
		return false
	}

	// Проверяем что есть точка в доменной части
	if !strings.Contains(parts[1], ".") {
		return false
	}

	return true
}

func isValidName(name string) bool {
	name = strings.TrimSpace(name)
	if len(name) < MinNameLength || len(name) > MaxNameLength {
		return false
	}
	return true
}

func isValidPassword(password string) bool {
	if len(password) < MinPasswordLenght || len(password) > MaxPasswordLenght {
		return false
	}
	return true
}

func (f *RegisterUserForm) Validate() error {
	if !isValidEmail(f.Email) {
		return ErrBadRequest
	}
	if !isValidName(f.Name) {
		return ErrBadRequest
	}
	if !isValidPassword(f.Password) {
		return ErrBadRequest
	}
	if f.Password != f.ConfirmPassword {
		return ErrConflict
	}
	return nil
}

func (f *ChangeUserForm) Validate() error {
	if !isValidName(f.Name) {
		return ErrBadRequest
	}
	return nil
}

func (f *ChangePassForm) Validate() error {
	if f.OldPassword == "" {
		return ErrBadRequest
	}
	if !isValidPassword(f.Password) {
		return ErrBadRequest
	}
	if f.Password != f.ConfirmPass {
		return ErrConflict
	}
	if f.OldPassword == f.Password {
		return ErrConflict
	}
	return nil
}

func (f *UserAuthForm) Validate() error {
	if f.Email == "" || f.Password == "" {
		return ErrBadRequest
	}
	if !isValidEmail(f.Email) {
		return ErrBadRequest
	}
	return nil
}
