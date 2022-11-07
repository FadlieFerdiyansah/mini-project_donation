package controllers

import (
	"fmt"
	"mini_project/config"
	"mini_project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var prog []models.Program

	if err := config.InitDB().Find(&prog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    prog,
	})
}

func Create(c echo.Context) error {
	userId, _ := strconv.Atoi(c.FormValue("user_id"))
	categoryId, _ := strconv.Atoi(c.FormValue("category_id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	expired_date := c.FormValue("expired_date")
	expectedAmount, _ := strconv.Atoi(c.FormValue("expected_amount"))
	fmt.Println(userId, categoryId, title, body, expectedAmount)
	prog := &models.Program{
		UserId:         uint64(userId),
		CategoryId:     uint64(categoryId),
		Title:          title,
		Body:           body,
		ExpiredDate:    expired_date,
		ExpectedAmount: int64(expectedAmount),
	}
	config.InitDB().Create(prog)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    prog,
	})
}

func Update(c echo.Context) error {
	userId, _ := strconv.Atoi(c.FormValue("user_id"))
	categoryId, _ := strconv.Atoi(c.FormValue("category_id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	expired_date := c.FormValue("expired_date")
	expectedAmount, _ := strconv.Atoi(c.FormValue("expected_amount"))
	fmt.Println(userId, categoryId, title, body, expectedAmount)
	prog := &models.Program{
		UserId:         uint64(userId),
		CategoryId:     uint64(categoryId),
		Title:          title,
		Body:           body,
		ExpiredDate:    expired_date,
		ExpectedAmount: int64(expectedAmount),
	}
	config.InitDB().Create(prog)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    prog,
	})
}
