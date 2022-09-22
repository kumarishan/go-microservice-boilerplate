package products

import (
	"context"
	"fmt"

	"github.com/kumarishan/errors"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
)

// Service defines interface for project service
type Service interface {
	GetProduct(ctx context.Context, id string) (*Product, error)
	AddProduct(ctx context.Context, name string) (*Product, error)
}

var _ = di.Provide(NewService)

type service struct {
	logger *logger.Logger
	repo   Repository
}

func NewService(logger *logger.Logger, repo Repository) Service {
	return &service{
		logger: logger,
		repo:   repo,
	}
}

func (s *service) GetProduct(ctx context.Context, id string) (*Product, error) {
	if err := ValidateProductId(id); err != nil {
		return nil, err
	}

	product, err := s.repo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			return nil, errors.Return(ErrProductNotFound, err, "")
		}

		return nil, errors.Return(errors.ErrInternal, err, "some internal error occured. please try again.")
	}

	return product, nil
}

// AddProduct implements Service
func (s *service) AddProduct(ctx context.Context, name string) (*Product, error) {
	if err := ValidateProductName(name); err != nil {
		return nil, errors.Return(errors.ErrInvalidInput, err, err.Error())
	}

	product := NewProduct(name)

	if _, err := s.repo.Save(ctx, product); err != nil {
		return nil, errors.Return(ErrAddingNewProduct, err, fmt.Sprintf("error adding new product %s", name))

	}
	return product, nil
}
