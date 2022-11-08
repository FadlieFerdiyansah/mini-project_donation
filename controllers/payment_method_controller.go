package controllers

import (
	"mini_project/config"
	"mini_project/helpers"
	"mini_project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexPayment(c echo.Context) error {
	var payments []models.PaymentMethod

	if err := config.InitDB().Find(&payments).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Payment method retrieved successfully", payments))
}

func CreatePayment(c echo.Context) error {
	name := c.FormValue("name")
	payment_code := c.FormValue("payment_code")

	payment := &models.PaymentMethod{
		Name:        name,
		PaymentCode: payment_code,
	}
	config.InitDB().Create(payment)

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Payment method created", payment))
}
