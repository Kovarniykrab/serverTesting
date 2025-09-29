package database

import (
	"context"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (rep *Repository) RegisterUser(ctx context.Context, form domain.RegisterUserForm) (domain.User, error) {
	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    form.Password,
	}

	_, err := rep.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		rep.log.Error("Failed to register user", "error", err)
		return domain.User{}, err
	}
	return user, err
}

func (rep *Repository) DeleteUser(ctx context.Context, id int) error {
	a, err := rep.db.NewDelete().Model(&domain.User{ID: id}).
		WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	rep.log.Debug("Delete", "deteled user", a)
	return err
}

func (rep *Repository) ChangeUser(ctx context.Context, id int, form domain.ChangeUserForm) (domain.ChangeUserForm, error) {
	user := domain.User{
		ID:          id,
		Name:        form.Name,
		DateOfBirth: form.DateOfBirth,
	}

	_, err := rep.db.NewUpdate().
		Model(&user).
		Where("id = ?", id).
		Exec(ctx)
	return form, err
}

func (rep *Repository) ChangePassword(ctx context.Context, id int, form domain.ChangePassForm) (domain.ChangePassForm, error) {
	_, err := rep.db.NewUpdate().
		Model(&domain.User{
			ID:       id,
			Password: form.Password,
		}).
		Column("password").
		Where("id = ?", id).
		Exec(ctx)
	return form, err
}

func (rep *Repository) GetUser(ctx context.Context, id int) (domain.User, error) {
	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		return user, err
	}

	return user, nil
}

func (rep *Repository) GetUserByEmail(ctx context.Context, Email string) (domain.User, error) {
	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("Email = ?", Email).Scan(ctx); err != nil {
		return user, err
	}

	return user, nil
}
