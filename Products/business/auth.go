package business

import (
	"log"

	"example.com/db"
	"example.com/structs"
)

type AuthBusiness struct {
	Postgres *db.DB
}

func NewAuthBusiness() *AuthBusiness {
	return &AuthBusiness{}
}

func (b *AuthBusiness) Connect(dbConnection *db.DB) error {
	b.Postgres = dbConnection
	return nil
}

func (p *AuthBusiness) GET() (interface{}, error) {

	return &structs.Auth{Email: "hassanalgo12@gmail.com", Password: "12345"}, nil
}

func (p *AuthBusiness) POST(Auth interface{}) error {
	return nil
}
func (p *AuthBusiness) Authenticate(Auth interface{}) error {
	log.Println("authenticated")
	log.Println(Auth)
	return nil
}
