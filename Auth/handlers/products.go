package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/structs"
	"github.com/labstack/echo/v4"
)

type ProductHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewProductHandler() *ProductHandlers {
	return &ProductHandlers{}
}

func (h *ProductHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

func (p *ProductHandlers) GET(ctx echo.Context) error {
	myData, _ := p.apiBusiness.GET()
	return ctx.JSON(http.StatusOK, myData)
}

func (p *ProductHandlers) POST(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get product")
}

func (p *ProductHandlers) PUT(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get product")
}
func (p *ProductHandlers) DELETE(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get product")
}

func (p *ProductHandlers) GETBYID(ctx echo.Context) error {
	id := ctx.Param("id")
	product, error := p.apiBusiness.GETBYID(structs.Product{ProductId: id})
	if error != nil {
		return ctx.JSON(http.StatusOK, error.Error())
	}
	return ctx.JSON(http.StatusOK, product)
}

func (p *ProductHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get product")
}
