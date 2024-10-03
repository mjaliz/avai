package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, from home!")
}
