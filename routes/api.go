package routes

import (
	"mini_project/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.POST("/login", controllers.Login)
	e.GET("/generate-hash", controllers.GenerateHashPassword)
	p := e.Group("programs")
	p.GET("", controllers.Index)
	p.POST("/create", controllers.Create)

	// e.GET("/about", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"status":  "200",
	// 		"message": "Successfully get the data",
	// 		"data":    "fadlie",
	// 	})
	// })

	return e
}
