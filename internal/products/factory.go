package products

import "github.com/rs/xid"

func NewProductId() ProductId {
	return ProductId(xid.New().String())
}

func NewProduct(name string) *Product {
	id := NewProductId()

	return &Product{
		ID:   id,
		Name: name,
	}
}
