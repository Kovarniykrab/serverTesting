package service

import (
	"context"
	"errors"
	"time"

	"github.com/Kovarniykrab/serverTesting/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	errorTimiout     = errors.New("Timeout")
	errorForbidden   = errors.New("Forbidden")
	errorConflict    = errors.New("Conflict")
	errorBadReq      = errors.New("Bad Request")
	errorUnautorized = errors.New("Unautorized")
	errorNoContent   = errors.New("No content")
	errorNoFound     = errors.New("No found")
	errorInternal    = errors.New("Internal")
	errorEmailOrPass = errors.New("email and password invalid")
	errorInvalid     = errors.New("Invalid")
)

func (app *Service) Register(ctx context.Context, form domain.RegisterUserForm) error {

	busyEmail, err := app.repositopy.GetUserByEmail(ctx, form.Email)
	if err == nil && busyEmail.ID != 0 {
		return domain.Conflict(errorConflict)
	}

	if form.Password != form.ConfirmPassword {
		return domain.BadRequest(errorBadReq)
	}

	hashPassword, e := app.Hash(form.Password)
	if e != nil {
		return errorInternal
		// ошибка, передать переменную ошибки.
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
		return domain.UserRender{}, "", domain.BadRequest(errorEmailOrPass)
	}

	user, err := app.repositopy.GetUserByEmail(ctx, form.Email)
	if err != nil {
		switch {
		case errors.Is(err, domain.NotFound(errorNoFound)):
			return domain.UserRender{}, "", domain.Unauthorized(errorEmailOrPass)
		case errors.Is(err, domain.BadRequest(errorBadReq)):
			return domain.UserRender{}, "", domain.BadRequest(errorInvalid)
		default:
			app.logger.Error("database error in Auth", "error", err, "email", form.Email)
			return domain.UserRender{}, "", domain.Internal(errorInternal)
		}
	}

	if err := Compare(user.Password, form.Password); err != nil {
		return domain.UserRender{}, "", domain.Unauthorized(errorEmailOrPass)
	}

	token, err := app.JWTService.CreateJWTToken(app.cfg.JWT, user.ID)
	if err != nil {
		return domain.UserRender{}, "", domain.Internal(errorInternal)
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
		case errors.Is(err, domain.NotFound(errorNoFound)):
			return domain.NotFound(errors.New("иди нахуй"))
		case errors.Is(err, domain.BadRequest(errorBadReq)):
			return domain.BadRequest(errorInvalid)
		default:
			app.logger.Error("database error in Delete", "error", err, "userID", id)
			return domain.Internal(errors.New("failed to delete user"))
		}
	}

	if err := app.repositopy.DeleteUser(ctx, id); err != nil {
		app.logger.Error("Failed to delete user", "error", err, "userID", id)
		return domain.Internal(errors.New("failed to delete user"))
	}

	app.logger.Info("User deleted successfully", "userID", id)
	return nil
}

func (app *Service) Update(ctx context.Context, id int, form domain.ChangeUserForm) error {

	user, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		return domain.Conflict(errorConflict)
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
		return domain.BadRequest(errors.New("invalid user id"))
	}
	if form.Password == "" || form.ConfirmPass == "" {
		return domain.BadRequest(errors.New("password fields cannot be empty"))
	}

	if form.Password != form.ConfirmPass {
		return domain.BadRequest(errors.New("passwords do not match"))
	}

	user, err := app.repositopy.GetUserById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return domain.NotFound(errorNoFound)

		case errors.Is(err, domain.ErrBadRequest):
			return domain.BadRequest(errorBadReq)

		case errors.Is(err, domain.ErrTimeout):
			return domain.Timeout(errorTimiout)

		default:

			app.logger.Error("database error in UpdatePassword", "error", err, "userID", id)
			return domain.Internal(errors.New("internal server error"))
		}
	}
	if err = Compare(user.Password, form.OldPassword); err != nil {
		return domain.Conflict(errorConflict)
	}

	if form.Password != form.ConfirmPass {
		return domain.Conflict(errorConflict)
	}

	hashPassword, err := app.Hash(form.Password)
	if err != nil {
		return domain.Conflict(errorConflict)
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
		return domain.UserRender{}, domain.NotFound(errorNoFound)
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
		return "", domain.Internal(errorInternal)

	}

	return string(hashed), err
}

func Compare(data string, aim string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data), []byte(aim))
	if err != nil {
		return domain.Forbidden(errorForbidden)
	}

	return nil
}
