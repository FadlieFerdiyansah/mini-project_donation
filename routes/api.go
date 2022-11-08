package routes

import (
	"mini_project/controllers"
	"mini_project/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	// e.GET("/generate-hash", controllers.GenerateHashPassword)

	p := e.Group("programs")
	p.GET("", controllers.Index)
	p.POST("/create", controllers.Create, middleware.IsAuthenticated)
	p.GET("/:id", controllers.Show)
	p.PUT("/:id/update", controllers.Update, middleware.IsAuthenticated)
	p.DELETE("/:id/destroy", controllers.Destroy, middleware.IsAuthenticated)
	p.POST("/:id/donation", controllers.Donation)

	pm := e.Group("payment-method")
	pm.GET("", controllers.IndexPayment)
	pm.POST("/create", controllers.CreatePayment, middleware.IsAuthenticated)

	return e
}
