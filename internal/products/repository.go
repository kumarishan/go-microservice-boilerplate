package products

import (
	"context"

	"github.com/kumarishan/errors"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
	"gorm.io/gorm"

	"go.uber.org/dig"
)

type Repository interface {
	repo.CrudRepository[Product, ProductId]

	FindByName(ctx context.Context, name string) ([]Product, error)
}

var _ = di.Provide(NewRepo)

type repository struct {
	*repo.CrudRepositoryImpl[Product, ProductId]
	db *gorm.DB
}

type RepoParams struct {
	dig.In
	Db *gorm.DB
}

func NewRepo(params RepoParams) Repository {
	impl := repo.NewCrudRepositoryImpl[Product, ProductId](params.Db)
	return &repository{
		impl,
		params.Db,
	}
}

func (r *repository) FindByName(ctx context.Context, name string) ([]Product, error) {
	var products []Product
	err := r.db.WithContext(ctx).Where("name = ?", name).Find(&products).Error
	if err != nil {
		return nil, errors.Return(err, nil, "")
	}
	return products, nil
}
