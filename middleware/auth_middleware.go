package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWT([]byte("miniproject"))
