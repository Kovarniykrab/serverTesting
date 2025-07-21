package domain

type UserForm struct {
	ID          int    `bun:"id" json:"id"`
	Email       string `bun:"email" json:"email"`
	Login       string `bun:"login" json:"login"`
	Password    string `bun:"password" json:"password"`
	PhoneNumber string `bun:"phone_number" json:"phone_number"`
}

type RegisterUserForm struct {
	Email           string `json:"email"`
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `validate:"required,eqfield=PasswordHash"`
	PhoneNumber     string `json:"phone_number"`
}

type UserAuthForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ChangePassForm struct {
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_pass" validate:"required,eqfield=Password"`
}

func (r *RegisterUserForm) ToUser() *UserForm {
	return &UserForm{
		Email:       r.Email,
		Login:       r.Login,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
	}
}
