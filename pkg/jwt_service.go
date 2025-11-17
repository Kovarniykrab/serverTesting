package pkg

import (
	"errors"
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

func (j *JWTService) CreateJWTToken(cnf configs.JWT, userID int) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cnf.HourExpired))),
		Issuer:    cnf.Issuer,
		Subject:   fmt.Sprintf("%d", userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTService) ValidateJwt(tokenString string) (*jwt.RegisteredClaims, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("malformed token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not valid yet")
			}
		}
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// архитектура папок
// не app методы, а обертки
//
