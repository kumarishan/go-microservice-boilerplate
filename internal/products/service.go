package products

import (
	"context"

	"github.com/kumarishan/errors"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/repo"
)

// Service defines interface for project service
type Service interface {
	GetProduct(ctx context.Context, id string) (*Product, error)
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
	product, err := s.repo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			return nil, errors.Return(ErrProductNotFound, err, "")
		}

		return nil, errors.Return(errors.ErrInternal, err, "some internal error occured. please try again.")
	}

	return product, nil
}
