package products

import (
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Product struct {
	ID   string
	Name string
}

type ProductPO struct {
	repo.Model[string]
	Name string
}

func (p ProductPO) ToEntity() *Product {
	return nil
}

func (p ProductPO) FromEntity(entity *Product) interface{} {
	return nil
}

func (p *ProductPO) BeforeCreate(tx *gorm.DB) error {
	p.ID = xid.New().String()

	return nil
}
