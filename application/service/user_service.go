package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/domain"
	"golang.org/x/crypto/bcrypt"
)

func (app *Service) RegisterUser(ctx context.Context, form domain.RegisterUserForm) error {

	busyEmail, err := app.re.GetUserByEmail(ctx, form.Email)
	if err == nil && busyEmail.ID != 0 {
		app.logger.Error("Email is busy", "error", form.Email)
		return fmt.Errorf("email уже занят")
	}

	if form.Password != form.ConfirmPassword {
		app.logger.Error("password don't match", "error", form.Password)
		return fmt.Errorf("пароли не совпадают")
	}

	hashPassword, e := app.Hash(form.Password)
	if e != nil {
		return e
	}

	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    hashPassword,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return app.re.RegisterUser(ctx, user)
}

func (app *Service) AuthUser(ctx context.Context, form domain.UserAuthForm) (domain.UserRender, error) {

	if form.Email == "" || form.Password == "" {
		return domain.UserRender{}, fmt.Errorf("логин и пароль обязательны")
	}

	user, err := app.re.GetUserByEmail(ctx, form.Email)
	if err != nil {
		return domain.UserRender{}, fmt.Errorf("пользователь не найден")
	}

	token, err := app.JWTService.CreateJWTToken(configs.JWT{}, user.ID)
	if err != nil {
		return domain.UserRender{}, fmt.Errorf("ошибка генерации токена")
	}

	return domain.UserRender{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}

func (app *Service) DeleteUser(ctx context.Context, id int) error {

	_, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}
	return app.re.DeleteUser(ctx, id)
}

func (app *Service) UpdateUser(ctx context.Context, id int, form domain.ChangeUserForm) error {

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		app.logger.Error("failed to get id", "error", err)
	}

	user.Name = form.Name
	user.DateOfBirth = form.DateOfBirth
	user.UpdatedAt = time.Now()

	return app.re.UpdateUser(ctx, user)
}

func (app *Service) UpdatePassword(ctx context.Context, id int, form domain.ChangePassForm) error {

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}

	if err = Compare(user.Password, form.OldPassword); err != nil {
		return fmt.Errorf("старый пароль не верен")
	}

	if form.Password != form.ConfirmPass {
		return fmt.Errorf("пароли не совпадают")
	}

	hashPassword, err := app.Hash(form.Password)
	if err != nil {
		return err
	}

	return app.re.ChangePassword(ctx, id, hashPassword)
}

func (app *Service) LogoutUser(ctx context.Context, id int) error {

	return nil
}

func (app *Service) GetUserById(ctx context.Context, id int) (user domain.User, err error) {

	return app.re.GetUserById(ctx, id)
}

func (s *Service) CheckUser(ctx context.Context, userID int) (domain.UserRender, error) {
	user, err := s.re.GetUserById(ctx, userID)
	if err != nil {
		return domain.UserRender{}, fmt.Errorf("пользователь не найден")
	}

	return domain.UserRender{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (app *Service) Hash(data string) (hash string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("ошибка хеширования %v", err)

		return
	}

	return string(hashed), err
}

func Compare(data string, aim string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data), []byte(aim))
	if err != nil {
		return err
	}

	return nil
}
