package handlers

import (
	"avai/internal/input"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func SignUp(c echo.Context) error {
	u := new(input.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := c.Validate(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	log.Println("Signing Up", u.Email)
	return nil
}
