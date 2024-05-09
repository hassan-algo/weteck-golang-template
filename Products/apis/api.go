package apis

import (
	"example.com/db"
	"github.com/labstack/echo/v4"
)

type APIBusiness interface {
	Connect(*db.DB) error
	GET() (interface{}, error)
	POST(interface{}) error
	MULTIPOST(interface{}) error
	PUT(interface{}) error
	GETBYID(interface{}) (interface{}, error)
	DELETE(interface{}) error
}

type APIHandler interface {
	Connect(APIBusiness) error
	GET(echo.Context) error
	POST(echo.Context) error
	MULTIPOST(echo.Context) error
	PUT(echo.Context) error
	DELETE(echo.Context) error
	GETBYID(echo.Context) error
}

type APIRouter interface {
	Connect(string, APIHandler, *echo.Echo, AuthHandler) error
}

type AuthBusiness interface {
	Connect(*db.DB) error
	GET() (interface{}, error)
	POST(interface{}) error
	Authenticate(interface{}) error
}

type AuthHandler interface {
	Connect(AuthBusiness) error
	GET(echo.Context) error
	POST(echo.Context) error
	Authenticate(func(ec echo.Context) error) func(ec echo.Context) error
}

type AuthRouter interface {
	Connect(string, AuthHandler, *echo.Echo) error
}
