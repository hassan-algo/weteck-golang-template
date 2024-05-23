package routes

import (
	"github.com/labstack/echo/v4"

	"example.com/apis"
)

type AuthRoutes struct {
}

func NewAuthRoutes() *AuthRoutes {
	return &AuthRoutes{}
}

func (r *AuthRoutes) Connect(endPoint string, AuthHandler apis.AuthHandler, echo *echo.Echo) error {
	echo.POST(endPoint, AuthHandler.Authentication)

	return nil
}
