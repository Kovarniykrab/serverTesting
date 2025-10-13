package service

import (
	"fmt"
	"time"

	"github.com/Kovarniykrab/serverTesting/configs"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secretKey string
}

func NewJWT(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (j *JWTService) CreateJWTToken(cnf configs.JWT, userID int) (*string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cnf.HourExpired))),
		Issuer:    cnf.Issuer,
		Subject:   fmt.Sprintf("%d", userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()), // Добавим время создания
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cnf.SecretKey))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (j *JWTService) ValidateJwt(tokenString string) (sk *jwt.StandardClaims, e error) {
	if tokenString == "" {
		return nil, e
	}

	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return sk, nil
	}

	// nolint:errorlint
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, err
		}

		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, err
		}

		return nil, err
	}

	return nil, err
}
