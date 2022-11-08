package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(pass string) string {
	hash, err := HashPassword(pass)
	if err != nil {
		panic(err)
	}
	return hash
}

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(pass, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false, nil
	}

	return true, nil
}
