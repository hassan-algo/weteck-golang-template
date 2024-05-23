filename=$(echo $1 | sed 's/\([A-Z]\)/_\L\1/g' | sed 's/^_\(.*\)/\1/')

touch business/${filename}.go
touch handlers/${filename}.go
touch routes/${filename}.go
touch structs/${filename}.go

cat << EOF > business/${filename}.go
package business

import (
	"example.com/db"
	"example.com/structs"
)

type ${1}Business struct {
	Postgres *db.DatabaseConnection
}

func New${1}Business() *${1}Business {
	return &${1}Business{}
}

func (b *${1}Business) Connect(dbConnection *db.DatabaseConnection) error {
	b.Postgres = dbConnection
	return nil
}

func (p *${1}Business) GET() (interface{}, error) {
	${filename}s := structs.${1}s{
		My${1}s: []structs.${1}{
		},
	}
	return ${filename}s, nil
}
func (p *${1}Business) GETBYID(data interface{}) (interface{}, error) {
	${filename}, _ := data.(structs.${1})
	return ${filename}, nil
}
func (p *${1}Business) POST(${filename} interface{}) error {
	return nil
}
func (p *${1}Business) MULTIPOST(${filename}s interface{}) error {
	return nil
}
func (p *${1}Business) PUT(${filename} interface{}) error {
	return nil
}
func (p *${1}Business) DELETE(${filename} interface{}) error {
	return nil
}
EOF


cat << EOF > handlers/${filename}.go
package handlers

import (
	"net/http"

	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type ${1}Handlers struct {
	apiBusiness apis.APIBusiness
}

func New${1}Handler() *${1}Handlers {
	return &${1}Handlers{}
}

func (h *${1}Handlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

func (p *${1}Handlers) GET(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET $1")
}

func (p *${1}Handlers) POST(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "POST $1")
}

func (p *${1}Handlers) PUT(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "PUT $1")
}
func (p *${1}Handlers) DELETE(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "DELETE $1")
}

func (p *${1}Handlers) GETBYID(ctx echo.Context) error {
	
	return ctx.JSON(http.StatusOK, "GETBYID $1")
}

func (p *${1}Handlers) MULTIPOST(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "MULTIPOST $1")
}
EOF


cat << EOF > routes/${filename}.go
package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type ${1}Routes struct {
}

func New${1}Routes() *${1}Routes {
	return &${1}Routes{}
}

func (r *${1}Routes) Connect(endPoint string, ${filename}Handler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, ${filename}Handler.GET)
	echo.POST(endPoint, ${filename}Handler.POST)
	echo.PUT(endPoint, ${filename}Handler.PUT)
	echo.DELETE(endPoint, ${filename}Handler.DELETE)
	echo.GET(endPoint+"/:id", ${filename}Handler.GETBYID)
	echo.POST(endPoint+"/multi", ${filename}Handler.MULTIPOST)
	return nil
}

EOF


cat << EOF > structs/${filename}.go
package structs

type ${1}s struct {
	My${1}s []${1} \`json:"${filename}s"\`
}

type ${1} struct {
	${1}Id   string \`json:"${filename}_id"\`
	${1}Name string \`json:"${filename}_name"\`
}

EOF



echo "Copy these lines in your main.go file"

echo "=========================================="
echo ""
echo "apis.NewAPI(\"/${filename}\","
echo "    postgres,"
echo "    routes.New${1}Routes(),"
echo "    handlers.New${1}Handler(),"
echo "    business.New${1}Business(), e, authAPI)"
echo ""
echo "=========================================="
