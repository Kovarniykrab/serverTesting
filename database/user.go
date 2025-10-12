package database

import (
	"context"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (rep *Repository) RegisterUser(ctx context.Context, form domain.User) error {

	_, err := rep.db.NewInsert().Model(&form).Exec(ctx)
	if err != nil {
		rep.log.Error("Failed to register user", "error", err)
		return err
	}
	return err
}

func (rep *Repository) DeleteUser(ctx context.Context, id int) error {

	a, err := rep.db.NewDelete().Model(&domain.User{ID: id}).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	rep.log.Debug("Delete", "deteled user", a)
	return err
}

func (rep *Repository) UpdateUser(ctx context.Context, user domain.User) error {

	_, err := rep.db.NewUpdate().Model(&user).WherePK().Exec(ctx)
	if err != nil {
		rep.log.Error("Failed to update user", "error", err)
		return err
	}
	return err
}

func (rep *Repository) ChangePassword(ctx context.Context, id int, hashedPassword string) error {

	_, err := rep.db.NewUpdate().
		Model(&domain.User{
			ID:        id,
			Password:  hashedPassword,
			UpdatedAt: time.Now(),
		}).
		Column("password", "updated_at").
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (rep *Repository) GetUserById(ctx context.Context, id int) (domain.User, error) {

	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		rep.log.Error("Failed to get user by id", "error", err)
		return domain.User{}, err
	}

	return user, nil
}

func (rep *Repository) GetUserByEmail(ctx context.Context, Email string) (domain.User, error) {

	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("Email = ?", Email).Scan(ctx); err != nil {
		rep.log.Error("Failed to get user by email", "error", err)
		return domain.User{}, err
	}

	return user, nil
}
