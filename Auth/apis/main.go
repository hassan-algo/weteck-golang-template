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

func NewAPI(endpoint string, connectionString string, routes APIRouter, handlers APIHandler, business APIBusiness, echo *echo.Echo) *API {
	newAPI := &API{
		ApiEndpoint: endpoint,
		ApiHandler:  handlers,
		ApiRoutes:   routes,
		ApiBusiness: business,
	}
	newAPI.ApiRoutes.Connect(newAPI.ApiEndpoint, newAPI.ApiHandler, echo)
	newAPI.ApiHandler.Connect(newAPI.ApiBusiness)
	newAPI.ApiBusiness.Connect(db.NewDBConnection(connectionString))
	return newAPI
}
