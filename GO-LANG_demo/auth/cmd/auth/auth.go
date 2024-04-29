package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	handlers "madeer.com/api/handlers"
	routes "madeer.com/api/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authHandle := handlers.NewAuthHandler(":LKdfa")
	authRoutes := routes.NewAuthRoutes(e, authHandle)
	authRoutes.AttachHandlers()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
