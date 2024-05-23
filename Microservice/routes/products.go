package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type ProductRoutes struct {
}

func NewProductRoutes() *ProductRoutes {
	return &ProductRoutes{}
}

func (r *ProductRoutes) Connect(endPoint string, productHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, auth.Authenticate(productHandler.GET))
	echo.POST(endPoint, productHandler.POST)
	echo.PUT(endPoint, productHandler.PUT)
	echo.DELETE(endPoint, productHandler.DELETE)
	echo.GET(endPoint+"/:id", productHandler.GETBYID)
	echo.POST(endPoint+"/multi", productHandler.MULTIPOST)
	return nil
}
