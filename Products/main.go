package main

import (
	"example.com/apis"
	"example.com/business"
	"example.com/db"
	"example.com/handlers"
	"example.com/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	postgres := db.NewDBConnection("connectionString")

	authAPI := apis.NewAUTH("/auth",
		postgres,
		routes.NewAuthRoutes(),
		handlers.NewAuthHandler(),
		business.NewAuthBusiness(), e)

	apis.NewAPI("/product",
		postgres,
		routes.NewProductRoutes(),
		handlers.NewProductHandler(),
		business.NewProductBusiness(), e, authAPI)
	apis.NewAPI("/products",
		postgres,
		routes.NewProductRoutes(),
		handlers.NewProductHandler(),
		business.NewProductBusiness(), e, authAPI)

	e.Start(":8080")

}
