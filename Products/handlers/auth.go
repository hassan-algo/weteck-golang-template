package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/structs"
	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	authBusiness apis.AuthBusiness
}

func NewAuthHandler() *AuthHandlers {
	return &AuthHandlers{}
}

func (h *AuthHandlers) Connect(business apis.AuthBusiness) error {
	h.authBusiness = business
	return nil
}

func (p *AuthHandlers) GET(ctx echo.Context) error {
	myData, _ := p.authBusiness.GET()
	return ctx.JSON(http.StatusOK, myData)
}

func (p *AuthHandlers) POST(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get Auth")
}
func (p *AuthHandlers) Authenticate(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		p.authBusiness.Authenticate(structs.Authenticate{Token: "qw34567890876564erthjk"})
		return f(ec)
	}
}
