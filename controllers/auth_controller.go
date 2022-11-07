package controllers

import (
	"mini_project/helpers"
	"mini_project/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// var pwd string
	email := c.FormValue("email")
	password := c.FormValue("password")
	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  400,
			"message": "No record found",
			"data":    nil,
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logged in",
		"data":    t,
	})
}

func GenerateHashPassword(c echo.Context) error {
	pass := c.FormValue("password")

	hash, _ := helpers.HashPassword(pass)
	return c.JSON(http.StatusOK, hash)
}
