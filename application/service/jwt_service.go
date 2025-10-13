package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secretKey string
}

func NewJWT(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (j *JWTService) GenerateJWT(userID int) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateJWT(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга токена: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, fmt.Errorf("невалидный токен")
}
