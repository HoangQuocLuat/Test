package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int64, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}