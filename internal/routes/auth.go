package routes

import (
	"github.com/labstack/echo/v4"
	"najva/internal/handlers"
)

func Auth(g *echo.Group) {
	g.POST("/sign-up", handlers.SignUp)
}
