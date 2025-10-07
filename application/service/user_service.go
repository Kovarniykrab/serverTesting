package service

import (
	"context"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (app *Service) RegisterUser(ctx context.Context, form domain.RegisterUserForm) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := domain.RegisterUserForm{
		DateOfBirth:     form.DateOfBirth,
		Name:            form.Name,
		Email:           form.Email,
		Password:        form.Password,
		ConfirmPassword: form.ConfirmPassword,
	}
	err := app.re.RegisterUser(ctx, user)
	return err
}

func (app *Service) AuthUser(ctx context.Context, form domain.UserAuthForm) (user domain.User, err error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return app.re.GetUserByEmail(ctx, form.Email)
}

func (app *Service) DeleteUser(ctx context.Context, id int) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return app.re.DeleteUser(ctx, id)
}

func (app *Service) UpdateUser(ctx context.Context, id int, form domain.ChangeUserForm) (domain.ChangeUserForm, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		app.logger.Error("failed to get id", "error", err)
	}
	if (user == domain.User{}) {
		app.logger.Error("user no found", "id", id)
	}

	return app.re.ChangeUser(ctx, id, form)
}

func (app *Service) UpdatePassword(ctx context.Context, id int, form domain.ChangePassForm) (domain.ChangePassForm, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return app.re.ChangePassword(ctx, id, form)
}

func (app *Service) LogoutUser(ctx context.Context, id int) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}

func (app *Service) GetUserById(ctx context.Context, id int) (user domain.User, err error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return app.re.GetUserById(ctx, id)
}
