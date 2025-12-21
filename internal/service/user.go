package service

import (
	"context"
	"errors"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
	"golang.org/x/crypto/bcrypt"
)

func (app *Service) Register(ctx context.Context, form domain.RegisterUserForm) error {

	busyEmail, err := app.repositopy.GetUserByEmail(ctx, form.Email)
	if err == nil && busyEmail.ID != 0 {
		return domain.ErrConflict
	}

	if form.Password != form.ConfirmPassword {
		return domain.ErrBadRequest
	}

	hashPassword, err := app.Hash(form.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    hashPassword,
		CreatedAt:   time.Now(),
	}
	return app.repositopy.UpdateUser(ctx, &user)
}

func (app *Service) Auth(ctx context.Context, form domain.UserAuthForm) (domain.UserRender, string, error) {

	if form.Email == "" || form.Password == "" {
		return domain.UserRender{}, "", domain.ErrBadRequest
	}

	user, err := app.repositopy.GetUserByEmail(ctx, form.Email)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return domain.UserRender{}, "", domain.ErrUnauthorized
		case errors.Is(err, domain.ErrBadRequest):
			return domain.UserRender{}, "", domain.ErrBadRequest
		default:
			app.logger.Error("database error in Auth", "error", err, "email", form.Email)
			return domain.UserRender{}, "", domain.ErrInternalServerError
		}
	}

	if err := Compare(user.Password, form.Password); err != nil {
		return domain.UserRender{}, "", domain.ErrUnauthorized
	}

	token, err := app.JWTService.CreateJWTToken(app.cfg.JWT, user.ID)
	if err != nil {
		return domain.UserRender{}, "", domain.ErrInternalServerError
	}

	userRender := domain.UserRender{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userRender, token, nil
}

func (app *Service) Delete(ctx context.Context, id int) error {

	_, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return domain.ErrNotFound
		case errors.Is(err, domain.ErrBadRequest):
			return domain.ErrBadRequest
		default:
			app.logger.Error("database error in Delete", "error", err, "userID", id)
			return domain.ErrInternalServerError
		}
	}

	if err := app.repositopy.DeleteUser(ctx, id); err != nil {
		app.logger.Error("Failed to delete user", "error", err, "userID", id)
		return domain.ErrInternalServerError
	}

	app.logger.Info("User deleted successfully", "userID", id)
	return nil
}

func (app *Service) Update(ctx context.Context, id int, form domain.ChangeUserForm) error {

	user, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		return domain.ErrConflict
	}

	user.Name = form.Name
	user.DateOfBirth = form.DateOfBirth
	user.UpdatedAt = Ptr(time.Now())

	return app.repositopy.UpdateUser(ctx, &user)
}

// func Ptr[T any](in T) *T {
//	return &in
//} что такое дженерики (прочитать)

func Ptr[T any](in T) *T {
	return &in
}

func (app *Service) UpdatePassword(ctx context.Context, id int, form domain.ChangePassForm) error {

	if id <= 0 {
		return domain.ErrBadRequest
	}
	if form.Password == "" || form.ConfirmPass == "" {
		return domain.ErrBadRequest
	}

	if form.Password != form.ConfirmPass {
		return domain.ErrBadRequest
	}

	user, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return domain.ErrNotFound

		case errors.Is(err, domain.ErrBadRequest):
			return domain.ErrBadRequest

		case errors.Is(err, domain.ErrTimeout):
			return domain.ErrTimeout

		default:

			app.logger.Error("database error in UpdatePassword", "error", err, "userID", id)
			return domain.ErrInternalServerError
		}
	}
	if err = Compare(user.Password, form.OldPassword); err != nil {
		return domain.ErrBadRequest
	}

	if form.Password != form.ConfirmPass {
		return domain.ErrBadRequest
	}

	hashPassword, err := app.Hash(form.Password)
	if err != nil {
		return domain.ErrBadRequest
	}

	user.Password = hashPassword

	return app.repositopy.UpdateUser(ctx, &user)
}

func (app *Service) GetById(ctx context.Context, id int) (userRender domain.UserRender, err error) {
	user, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		return domain.UserRender{}, err
	}
	userRender = domain.UserRender{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userRender, nil
}

func (s *Service) Check(ctx context.Context, userID int) (domain.UserRender, error) {
	user, err := s.repositopy.GetUserById(ctx, userID)
	if err != nil {
		return domain.UserRender{}, domain.ErrNotFound
		// ветвление ошибок
	}

	return domain.UserRender{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (app *Service) Hash(data string) (hash string, err error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrInternalServerError

	}

	return string(hashed), err
}

func Compare(data string, aim string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data), []byte(aim))
	if err != nil {
		return domain.ErrForbidden
	}

	return nil
}
