package products

import (
	"context"

	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
	"gorm.io/gorm"

	"go.uber.org/dig"
)

type Repository interface {
	repo.CrudRepository[ProductPO, Product, string]

	FindByName(ctx context.Context, name string) ([]Product, error)
}

var _ = di.Provide(NewRepo)

type repository struct {
	*repo.CrudRepositoryImpl[ProductPO, Product, string]
	db *gorm.DB
}

type RepoParams struct {
	dig.In
	db *gorm.DB
}

func NewRepo(params RepoParams) Repository {
	impl := repo.NewCrudRepositoryImpl[ProductPO, Product, string](params.db)
	return &repository{
		impl,
		params.db,
	}
}

func (r *repository) FindByName(ctx context.Context, name string) ([]Product, error) {
	var products []Product
	err := r.db.WithContext(ctx).Where("name = ?", name).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
