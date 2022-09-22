package orders

import (
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
)

type OrderId string

func (oi OrderId) String() string {
	return string(oi)
}

type Order struct {
	repo.Model[OrderId]
	Items []products.Product
}
