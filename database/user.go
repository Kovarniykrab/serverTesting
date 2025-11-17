package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Kovarniykrab/serverTesting/domain"
)

func (rep *Repository) DeleteUser(ctx context.Context, id int) error {

	user, er := rep.GetUserById(ctx, id)
	if er != nil {
		rep.log.Error("User not found for deletion", "id", id, "error", er)
		return er
	}

	if user.ID <= 0 {
		return domain.BadRequest(errors.New("invalid user id"))
	}

	_, err := rep.db.NewDelete().Model(&domain.User{ID: id}).WherePK().Exec(ctx)
	if err != nil {
		rep.log.Error("Failed to delete user", "error", err, "userID", id)
		return domain.Internal(errors.New("failed to delete user"))
	}

	rep.log.Debug("Delete", "deteled user", id)
	return nil
}

func (rep *Repository) UpdateUser(ctx context.Context, user *domain.User) error {
	if user.ID == 0 {
		_, err := rep.db.NewInsert().Model(user).Exec(ctx)
		return err
	} else {
		_, err := rep.db.NewInsert().Model(user).On("CONFLICT (id) DO UPDATE").Exec(ctx)
		return err
	}
}

func (rep *Repository) GetUserById(ctx context.Context, id int) (domain.User, error) {

	if id <= 0 {
		return domain.User{}, domain.BadRequest(errors.New("invalid user id"))
	}

	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		rep.log.Error("Failed to get user by id", "error", err)
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.NotFound(errors.New("user not found"))
		}

		if isConnectionError(err) {
			return user, domain.Internal(errors.New("database connection failed"))
		}

		if isTimeoutError(err) {
			return user, domain.Internal(errors.New("database timeout"))
		}

		return user, domain.Internal(errors.New("failed to get user"))
	}
	// ошибка поиска по id. Отлавливать и возвращать самостоятельно
	return user, nil
}

func (rep *Repository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {

	if email == "" {
		return domain.User{}, domain.BadRequest(errors.New("email cannot be empty"))
	}

	user := domain.User{}

	if err := rep.db.NewSelect().Model(&user).Where("Email = ?", email).Scan(ctx); err != nil {
		rep.log.Error("Failed to get user by email", "error", err)
		rep.log.Error("Failed to get user by email", "error", err, "email", email)

		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.NotFound(errors.New("user not found"))
		}

		if isConnectionError(err) {
			return domain.User{}, domain.Internal(errors.New("database connection failed"))
		}

		return domain.User{}, domain.Internal(errors.New("failed to get user"))
	}

	return user, nil
}

// func GetAll

func isConnectionError(err error) bool {
	errorText := err.Error()
	return strings.Contains(errorText, "connection") ||
		strings.Contains(errorText, "network") ||
		strings.Contains(errorText, "refused") ||
		strings.Contains(errorText, "unreachable")
}

func isTimeoutError(err error) bool {
	errorText := err.Error()
	return strings.Contains(errorText, "timeout") ||
		strings.Contains(errorText, "deadline") ||
		strings.Contains(errorText, "context canceled")
}

// уровни мидлвар + корсы
// хардкод убрать
// функция search с фильтрами по имени, датами создания и удаления, по дате рождения, по полу
// квери запросы
