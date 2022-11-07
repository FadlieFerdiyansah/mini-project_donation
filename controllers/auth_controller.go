package controllers

import (
	"mini_project/helpers"
	"mini_project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, "No record found", nil))
	}

	if !res {
		return c.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, "There's something went wrong", nil))
	}

	token, err := helpers.CreateToken(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Login successfully", token))
}

func GenerateHashPassword(c echo.Context) error {
	pass := c.FormValue("password")

	hash, _ := helpers.HashPassword(pass)
	return c.JSON(http.StatusOK, hash)
}
