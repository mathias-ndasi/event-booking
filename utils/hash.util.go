package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), error
}

func IsPasswordHashValid(password string, hashedPassword string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return error == nil
}
