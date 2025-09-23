package service

import (
	"context"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (app *App) RegisterUser(ctx context.Context, form domain.RegisterUserForm) (domain.User, error) {
	return app.Database.RegisterUser(ctx, form)
}

func (app *App) AuthUser(ctx context.Context, form domain.UserAuthForm) (user domain.User, err error) {
	return app.Database.ChangeUser(ctx, form)
}

func (app *App) DeleteUser(ctx context.Context, id int) (domain.User, error) {
	return app.Database.GetUser(ctx, id)

}

func (app *App) UpdatePassword(ctx context.Context, id int) (user domain.ChangeUserForm, err error) {
	return app.Database.ChangeUser()
}

func (app *App) LogoutUser(ctx context.Context, id int) error {
	return app.Database.ChangeUser(ctx, &domain.User{})
}

func (app *App) GetUser(ctx context.Context, id int) (user domain.User, err error) {
	return app.Database.GetUser()
}
