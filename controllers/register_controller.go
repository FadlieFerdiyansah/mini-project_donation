package controllers

import (
	"mini_project/config"
	"mini_project/helpers"
	"mini_project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	phone := c.FormValue("phone")

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: helpers.GenerateHashPassword(password),
		Phone:    phone,
	}
	config.InitDB().Create(user)

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "User created", user))
}
