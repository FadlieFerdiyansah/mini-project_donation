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

	return e
}
