package apis

import (
	"example.com/db"
	"github.com/labstack/echo/v4"
)

type API struct {
	ApiHandler  APIHandler
	ApiEndpoint string
	ApiRoutes   APIRouter
	ApiBusiness APIBusiness
}

type AUTH struct {
	ApiHandler  AuthHandler
	ApiEndpoint string
	ApiRoutes   AuthRouter
	ApiBusiness AuthBusiness
}

func NewAPI(endpoint string, postgres *db.DB, routes APIRouter, handlers APIHandler, business APIBusiness, echo *echo.Echo, authAPI *AUTH) *API {
	newAPI := &API{
		ApiEndpoint: endpoint,
		ApiHandler:  handlers,
		ApiRoutes:   routes,
		ApiBusiness: business,
	}
	newAPI.ApiRoutes.Connect(newAPI.ApiEndpoint, newAPI.ApiHandler, echo, authAPI.ApiHandler)
	newAPI.ApiHandler.Connect(newAPI.ApiBusiness)
	newAPI.ApiBusiness.Connect(postgres)
	return newAPI
}

func NewAUTH(endpoint string, postgres *db.DB, routes AuthRouter, handlers AuthHandler, business AuthBusiness, echo *echo.Echo) *AUTH {
	newAPI := &AUTH{
		ApiEndpoint: endpoint,
		ApiHandler:  handlers,
		ApiRoutes:   routes,
		ApiBusiness: business,
	}
	newAPI.ApiRoutes.Connect(newAPI.ApiEndpoint, newAPI.ApiHandler, echo)
	newAPI.ApiHandler.Connect(newAPI.ApiBusiness)
	newAPI.ApiBusiness.Connect(postgres)
	return newAPI
}
