package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Default(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}