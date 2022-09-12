package products

import (
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
	"gorm.io/gorm"
)

type productPM struct {
	repo.Model[ProductId]
	Name string
}

func (productPM) TableName() string {
	return "products"
}

func (p productPM) ToEntity() *Product {
	return &Product{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p productPM) FromEntity(entity *Product) interface{} {
	return &productPM{
		Model: repo.Model[ProductId]{
			ID: entity.ID,
		},
		Name: entity.Name,
	}
}

func (p *productPM) BeforeCreate(tx *gorm.DB) error {
	return nil
}
