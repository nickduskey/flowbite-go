package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nickduskey/flowbite-go/internal/pkg/handlers"
)

func main() {
	app := echo.New()

	// app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	app.Static("/", "assets")
	app.Static("/flowbite", "node_modules/flowbite/dist")
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	handlers.SetupRoutes(app)

	app.Logger.Fatal(app.Start(":8080"))
}
