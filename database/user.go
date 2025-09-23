package database

import (
	"context"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (db *Service) RegisterUser(ctx context.Context, form domain.RegisterUserForm) (domain.User, error) {
	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    form.Password,
	}

	_, err := db.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		db.log.Error("Failed to register user", "error", err)
		return domain.User{}, err
	}
	return user, nil
}

func (db *Service) DeleteUser(ctx context.Context, id int) error {
	a, err := db.db.NewDelete().Model(&domain.User{ID: id}).
		WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	db.log.Debug("Delete", "deteled user", a)
	return nil
}

func (db *Service) ChangeUser(ctx context.Context, form domain.ChangeUserForm) (domain.ChangeUserForm, error) {
	if _, err := db.db.NewInsert().
		Model(&form).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx); err != nil {
		return form, err
	}

	return form, nil
}

func (db *Service) GetUser(ctx context.Context, id int) (domain.User, error) {
	md := domain.User{}

	if err := db.db.NewSelect().Model(&md).Where("id = ?", id).Scan(ctx); err != nil {
		return md, err
	}

	return md, nil
}
