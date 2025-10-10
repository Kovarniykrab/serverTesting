package domain

import (
	"time"

	"github.com/uptrace/bun"
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

// форма смены данных пользователя
type ChangeUserForm struct {
	//имя пользователя
	// Поле обязательно
	Name string `bun:"name" json:"name" binding:"required"`
	//дата рождения
	// Поле обязательно
	DateOfBirth string `bun:"date_of_birth" json:"date_of_birth" binding:"required"`

	UpdatedAt string `bun:"updated_at" json:"updated_at" binding:"required"`
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
	UpdatedAt   string `bun:"updated_at" json:"updated_at" binding:"required"`
}

type Users []User
