package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"

	"example.com/apis"
	"example.com/business"
	"example.com/db"
	"example.com/handlers"
	"example.com/routes"
)

func main() {

	e := echo.New()

	postgres := db.NewDatabaseConnection()

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

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	go func() {
		e.Start(":8081")
	}()

	<-sigChannel
	postgres.Con.Close()
	fmt.Println("Database connection closed!")

}
