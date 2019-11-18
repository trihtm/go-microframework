package internal

import (
	"github.com/labstack/echo"
	"go-microframework/internal/handler"
	"go-microframework/pkg/database"
)

func RegisterRoutes(e *echo.Echo, databaseManager *database.Manager) {
	e.GET("/", handler.Default)
}