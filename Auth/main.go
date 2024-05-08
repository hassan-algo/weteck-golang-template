package main

import (
	"example.com/apis"
	"example.com/business"
	"example.com/handlers"
	"example.com/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	apis.NewAPI("/product",
		"connectionString",
		routes.NewProductRoutes(),
		handlers.NewProductHandler(),
		business.NewProductBusiness(), e)

	e.Start(":8080")

	// fmt.Println("Starting")
	// var product apis.APIHandler = &modules.Product{
	// 	ProductId:   "abc",
	// 	ProductName: "xyz",
	// }
	// if p, err := product.GET(); err == nil {
	// 	fmt.Println("Product", p)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// product = &modules.Product{
	// 	ProductId: "123872931",
	// }

	// fmt.Println("Product by id")
	// if err := product.GETBYID(); err == nil {
	// 	fmt.Println("Product new", product)
	// } else {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Product by Post method")
	// if err := product.POST(); err == nil {
	// 	fmt.Println("Product", product)
	// } else {
	// 	fmt.Println(err.Error())
	// }

}
