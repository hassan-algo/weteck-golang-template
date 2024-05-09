package structs

type Products struct {
	MyProducts []Product `json:"products"`
}

type Product struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
