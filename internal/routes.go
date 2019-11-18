package internal

import (
	"github.com/labstack/echo"
	"go-microframework/internal/handler"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handler.Default)
}