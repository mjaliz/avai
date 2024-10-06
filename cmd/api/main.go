package main

import (
	"avai/internal/app"
	"avai/internal/db"
	"avai/internal/handlers"
	"avai/internal/input"
	"avai/internal/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app.NewApp(db.InitDB())
	e := echo.New()
	e.Validator = &input.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	auth := e.Group("/auth")
	routes.Auth(auth)
	e.GET("/", handlers.Root)
	e.Logger.Fatal(e.Start(":8000"))
}
