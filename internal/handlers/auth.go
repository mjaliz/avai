package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
)

func SignUp(c echo.Context) error {
	log.Println("Signing Up")
	return nil
}
