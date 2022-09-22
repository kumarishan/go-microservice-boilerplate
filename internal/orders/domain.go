package orders

import "github.com/kumarishan/go-microservice-boilerplate/internal/products"

type OrderId string

func (oi OrderId) String() string {
	return string(oi)
}

type Order struct {
	ID    OrderId
	Items []products.Product
}
