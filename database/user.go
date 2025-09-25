package database

import (
	"context"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (db *Repository) RegisterUser(ctx context.Context, form domain.RegisterUserForm) (domain.User, error) {
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
	return user, err
}

func (db *Repository) DeleteUser(ctx context.Context, id int) error {
	a, err := db.db.NewDelete().Model(&domain.User{ID: id}).
		WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	db.log.Debug("Delete", "deteled user", a)
	return err
}

func (db *Repository) ChangeUser(ctx context.Context, form domain.ChangeUserForm) (domain.ChangeUserForm, error) {
	if _, err := db.db.NewInsert().
		Model(&form).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx); err != nil {
		return form, err
	}

	return form, nil
}

func (db *Repository) GetUser(ctx context.Context, id int) (domain.User, error) {
	user := domain.User{}

	if err := db.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		return user, err
	}

	return user, nil
}

func (db *Repository) GetUserByEmail(ctx context.Context, Email string) (domain.User, error) {
	user := domain.User{}

	if err := db.db.NewSelect().Model(&user).Where("Email = ?", Email).Scan(ctx); err != nil {
		return user, err
	}

	return user, nil
}
