package business

import (
	"errors"

	"example.com/db"
	"example.com/structs"
)

type ProductBusiness struct {
	Postgres *db.DB
}

func NewProductBusiness() *ProductBusiness {
	return &ProductBusiness{}
}

func (b *ProductBusiness) Connect(dbConnection *db.DB) error {
	b.Postgres = dbConnection
	return nil
}

func (p *ProductBusiness) GET() (interface{}, error) {
	products := structs.Products{
		MyProducts: []structs.Product{
			structs.Product{
				ProductId:   "123",
				ProductName: "Product name",
			},
			structs.Product{
				ProductId:   "123456",
				ProductName: "Product name2",
			},
		},
	}
	return products, nil
}
func (p *ProductBusiness) GETBYID(data interface{}) (interface{}, error) {
	product, ok := data.(structs.Product)
	if !ok {
		return nil, errors.New("Invalid product")
	}
	if product.ProductId != "1234" {
		return nil, errors.New("Product Not Found")
	}
	product.ProductName = "WOrking"
	return product, nil
}
func (p *ProductBusiness) POST(product interface{}) error {
	return nil
}
func (p *ProductBusiness) MULTIPOST(products interface{}) error {
	return nil
}
func (p *ProductBusiness) PUT(product interface{}) error {
	return nil
}
func (p *ProductBusiness) DELETE(product interface{}) error {
	return nil
}
