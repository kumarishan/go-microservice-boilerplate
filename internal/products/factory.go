package products

import (
	"time"

	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

func NewProductId() ProductId {
	return ProductId(xid.New().String())
}

func NewProduct(name string) *Product {
	id := NewProductId()

	return &Product{
		repo.Model[ProductId]{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		name,
		StatusActive,
	}
}
