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

func (p *productPM) MapToEntity() *Product {
	return &Product{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p *productPM) MapFromEntity(entity *Product) *productPM {
	p.Model = repo.Model[ProductId]{
		ID: entity.ID,
	}
	p.Name = entity.Name
	return p
}

func (p *productPM) BeforeCreate(tx *gorm.DB) error {
	return nil
}
