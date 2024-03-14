package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashpassword), nil
}

func VerifyPassword(password string, hashpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
}