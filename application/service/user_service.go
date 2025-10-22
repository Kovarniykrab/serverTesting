package service

import (
	"context"
	"time"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/Kovarniykrab/serverTesting/domain"
	"golang.org/x/crypto/bcrypt"
)

func (app *Service) RegisterUser(ctx context.Context, form domain.RegisterUserForm) error {

	busyEmail, err := app.re.GetUserByEmail(ctx, form.Email)
	if err == nil && busyEmail.ID != 0 {
		return domain.Conflict(err)
	}

	if form.Password != form.ConfirmPassword {
		return domain.BadRequest(err)
	}

	hashPassword, e := app.Hash(form.Password)
	if e != nil {
		return e
	}

	user := domain.User{
		DateOfBirth: form.DateOfBirth,
		Name:        form.Name,
		Email:       form.Email,
		Password:    hashPassword,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return app.re.RegisterUser(ctx, user)
}

func (app *Service) AuthUser(ctx context.Context, form domain.UserAuthForm) (domain.UserRender, error) {

	if form.Email == "" || form.Password == "" {
		var err error
		return domain.UserRender{}, domain.Unauthorized(err)
	}

	user, err := app.re.GetUserByEmail(ctx, form.Email)
	if err != nil {
		return domain.UserRender{}, domain.Unauthorized(err)
	}

	if err := Compare(user.Password, form.Password); err != nil {
		return domain.UserRender{}, domain.Unauthorized(err)
	}

	token, err := app.JWTService.CreateJWTToken(configs.JWT{}, user.ID)
	if err != nil {
		return domain.UserRender{}, domain.Unauthorized(err)
	}

	return domain.UserRender{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}

func (app *Service) DeleteUser(ctx context.Context, id int) error {

	_, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return domain.NotFound(err)
	}

	// проверка на емейл
	return app.re.DeleteUser(ctx, id)
}

func (app *Service) UpdateUser(ctx context.Context, id int, form domain.ChangeUserForm) error {

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return domain.Conflict(err)
	}

	user.Name = form.Name
	user.DateOfBirth = form.DateOfBirth
	user.UpdatedAt = time.Now()

	return app.re.UpdateUser(ctx, user)
}

func (app *Service) UpdatePassword(ctx context.Context, id int, form domain.ChangePassForm) error {

	user, err := app.re.GetUserById(ctx, id)
	if err != nil {
		return domain.NotFound(err)
	}
	// проверка на емайл
	if err = Compare(user.Password, form.OldPassword); err != nil {
		return domain.Conflict(err)
	}

	if form.Password != form.ConfirmPass {
		return domain.Conflict(err)
	}

	hashPassword, err := app.Hash(form.Password)
	if err != nil {
		return domain.Conflict(err)
	}

	return app.re.ChangePassword(ctx, id, hashPassword)
}

func (app *Service) LogoutUser(ctx context.Context, id int) error {

	return nil
}

func (app *Service) GetUserById(ctx context.Context, id int) (user domain.User, err error) {

	return app.re.GetUserById(ctx, id)
}

func (s *Service) CheckUser(ctx context.Context, userID int) (domain.UserRender, error) {
	user, err := s.re.GetUserById(ctx, userID)
	if err != nil {
		return domain.UserRender{}, domain.NotFound(err)
	}

	return domain.UserRender{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (app *Service) Hash(data string) (hash string, err error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.Internal(err)

	}

	return string(hashed), err
}

func Compare(data string, aim string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data), []byte(aim))
	if err != nil {
		return domain.Forbidden(err)
	}

	return nil
}
