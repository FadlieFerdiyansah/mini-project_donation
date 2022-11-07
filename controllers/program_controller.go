package controllers

import (
	"mini_project/config"
	"mini_project/helpers"
	"mini_project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var progams []models.Program

	if err := config.InitDB().Find(&progams).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Programs retrieved successfully", progams))
}

func Create(c echo.Context) error {
	userId, _ := strconv.Atoi(c.FormValue("user_id"))
	categoryId, _ := strconv.Atoi(c.FormValue("category_id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	expired_date := c.FormValue("expired_date")
	expectedAmount, _ := strconv.Atoi(c.FormValue("expected_amount"))

	prog := &models.Program{
		UserId:         uint64(userId),
		CategoryId:     uint64(categoryId),
		Title:          title,
		Body:           body,
		ExpiredDate:    expired_date,
		ExpectedAmount: int64(expectedAmount),
	}
	config.InitDB().Create(prog)

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Program created", prog))
}

func Update(c echo.Context) error {
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	expired_date := c.FormValue("expired_date")
	expected_amount, _ := strconv.Atoi(c.FormValue("expected_amount"))

	programId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(http.StatusNotFound, "Program not found", nil))
	}
	var program models.Program
	program.ID = uint64(programId)

	err = c.Bind(&program)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, err.Error(), nil))
	}

	con := config.InitDB()
	con.Model(&models.Program{}).Where("id = ?", program.ID).Updates(&models.Program{Title: title, Body: body, ExpiredDate: expired_date, ExpectedAmount: int64(expected_amount), CategoryId: uint64(category_id)})
	con.Find(&program, "id = ?", program.ID)
	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Successfully Update the program", program))
}

func Show(c echo.Context) error {
	programId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(http.StatusNotFound, "Program not found", nil))
	}
	var program models.Program
	con := config.InitDB()
	con.Find(&program, "id = ?", programId)

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "View program", program))
}

func Destroy(c echo.Context) error {
	programId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(http.StatusNotFound, "Program not found", nil))
	}
	var program models.Program
	con := config.InitDB()
	con.Delete(&program, programId)

	return c.JSON(http.StatusOK, helpers.Response(http.StatusOK, "Successfully deleted the program", nil))
}
