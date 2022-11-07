package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(email string) (string, error) {
	mapClaims := jwt.MapClaims{}
	mapClaims["email"] = email
	mapClaims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
