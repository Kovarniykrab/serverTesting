package domain

type RegisterUser struct {
	Email           string `bun:"email" json:"email"`
	Login           string `bun:"login" json:"login"`
	PasswordHash    string `bun:"password_hash" json:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=PasswordHash"`
	PhoneNumber     string `bun:"phone" json:"phone"`
}

type UserAuth struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type ChangePass struct {
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_pass" validate:"required,eqfield=Password"`
}
