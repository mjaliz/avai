package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"najva/internal/handlers"
	"najva/internal/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	auth := e.Group("/auth")
	routes.Auth(auth)
	e.GET("/", handlers.Root)
	e.Logger.Fatal(e.Start(":8000"))
}
