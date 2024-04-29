package routes

import (
	"github.com/labstack/echo"
	handlers "madeer.com/api/handlers"
)

type AuthRoutes struct {
	e        *echo.Echo
	handlers *handlers.AuthHanders
}

func NewAuthRoutes(ec *echo.Echo, handler *handlers.AuthHanders) *AuthRoutes {
	return &AuthRoutes{e: ec, handlers: handler}
}

func (ar *AuthRoutes) AttachHandlers() error {
	ar.e.GET("/", ar.handlers.Get)
	ar.e.POST("/", ar.handlers.Post)
	return nil
}
