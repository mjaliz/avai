package routes

import (
	"avai/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group) {
	g.POST("/sign-up", handlers.SignUp)
}
