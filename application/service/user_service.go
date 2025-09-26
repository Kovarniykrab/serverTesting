package service

import (
	"context"
	"errors"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (app *Service) RegisterUser(ctx context.Context, form domain.RegisterUserForm) error {
	if form.Email == "" || form.Password == "" || form.Name == "" || form.DateOfBirth == "" {
		return errors.New("все поля обязательны")
	}

	if form.Password != form.ConfirmPassword {
		return errors.New("пароли не совпадают")
	}

	existingUser, err := app.repo.GetUserByEmail(ctx, form.Email)
	if err == nil && existingUser.ID != 0 {
		return errors.New("email уже занят")
	}

	_, err = app.repo.RegisterUser(ctx, form)
	return err
}

func (app *Service) AuthUser(ctx context.Context, form domain.UserAuthForm) (user domain.User, err error) {
	return app.repo.GetUserByEmail(ctx, form.Email)
}

func (app *Service) DeleteUser(ctx context.Context, id int) error {
	return app.repo.DeleteUser(ctx, id)
}

func (app *Service) UpdateUser(ctx context.Context, id int, form domain.ChangeUserForm) error {
	return nil
}

func (app *Service) UpdatePassword(ctx context.Context, id int, form domain.ChangePassForm) error {
	return nil
}

func (app *Service) LogoutUser(ctx context.Context, id int) error {
	return nil
}

func (app *Service) GetUser(ctx context.Context, id int) (user domain.User, err error) {
	return app.repo.GetUser(ctx, id)
}
