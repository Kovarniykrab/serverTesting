package domain

type User struct {
	ID           int    `bun:"id" json:"id"`
	Email        string `bun:"email" json:"email"`
	Login        string `bun:"login" json:"login"`
	PasswordHash string `bun:"password_hash" json:"-"`
	PhoneNumber  string `bun:"phone" json:"phone"`
}

type RegisterUser struct {
	Email           string `json:"email"`
	Login           string `json:"login"`
	PasswordHash    string `json:"password"`
	ConfirmPassword string `validate:"required,eqfield=PasswordHash"`
	PhoneNumber     string `json:"phone"`
}

type UserAuth struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type ChangePass struct {
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_pass" validate:"required,eqfield=Password"`
}

func (r *RegisterUser) toUser() *User {
	return &User{
		Email:        r.Email,
		Login:        r.Login,
		PasswordHash: r.PasswordHash,
		PhoneNumber:  r.PhoneNumber,
	}
}
