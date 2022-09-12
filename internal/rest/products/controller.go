package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/errors"
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
)

var _ = di.Provide(NewProductsController)

// Controller is a products controller interface
type Controller interface {
	GetProduct(ctx *gin.Context) (int, *ProductDto, error)
	AddProduct(ctx *gin.Context) (int, *ProductDto, error)
}

// impl
type controller struct {
	log     *logger.Logger
	service products.Service
}

func NewProductsController(log *logger.Logger, service products.Service) Controller {
	return &controller{log, service}
}

func (p *controller) GetProduct(ctx *gin.Context) (int, *ProductDto, error) {
	id := ctx.Param("id")

	product, err := p.service.GetProduct(ctx, id)

	if err != nil {
		if errors.Is(err, errors.ErrInvalidInput) {
			return http.StatusBadRequest, nil, err
		}

		if errors.Is(err, products.ErrProductNotFound) {
			return http.StatusNotFound, nil, err
		}

		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &ProductDto{
		Id:   product.ID.String(),
		Name: product.Name,
	}, nil
}

// AddProduct implements Controller
func (p *controller) AddProduct(ctx *gin.Context) (int, *ProductDto, error) {

	var request AddProductRequest
	if err := ctx.BindJSON(&request); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := request.Validate(); err != nil {
		return http.StatusBadRequest, nil, err
	}

	product, err := p.service.AddProduct(ctx, request.Name)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &ProductDto{
		Id:   product.ID.String(),
		Name: product.Name,
	}, nil

}
