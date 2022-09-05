package products

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
)

var _ = di.Provide(NewProductsController)

type Controller interface {
	GetProduct(ctx *gin.Context) (int, *products.Product, error)
}

type controller struct {
	log     *logger.Logger
	service products.Service
}

func NewProductsController(log *logger.Logger, service products.Service) Controller {
	return &controller{log, service}
}

func (p *controller) GetProduct(ctx *gin.Context) (int, *products.Product, error) {
	id := ctx.Param("id")
	product, err := p.service.GetProduct(ctx, id)

	if err != nil {
		if errors.Is(err, products.ErrProductNotFound) {
			return http.StatusNotFound, nil, err
		}

		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, product, nil
}
