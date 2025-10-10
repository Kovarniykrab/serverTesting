package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
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

	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    form.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return app.re.RegisterUser(ctx, user)
}

func (app *Service) AuthUser(ctx context.Context, form domain.UserAuthForm) (user domain.User, err error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return app.re.GetUserByEmail(ctx, form.Email)
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

	if form.OldPassword != form.Password {
		return fmt.Errorf("пароли не совпадают")
	}

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}

	user.Password = form.Password
	user.UpdatedAt = time.Now()

	return app.re.ChangePassword(ctx, id, form)
}

func (app *Service) LogoutUser(ctx context.Context, id int) error {

	return nil
}

func (app *Service) GetUserById(ctx context.Context, id int) (user domain.User, err error) {

	return app.re.GetUserById(ctx, id)
}
