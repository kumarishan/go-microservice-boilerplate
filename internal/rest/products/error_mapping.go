package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
)

type ErrorResponse struct {
	Err            error
	HttpStatusCode int
	ErrorCode      string
	Message        string
	Args           []interface{}
}

type RouteErrorMaps struct {
	Mappings []ErrorResponse
}

func AddProduct(ctx *gin.Context) RouteErrorMaps {
	return RouteErrorMaps{
		Mappings: []ErrorResponse{
			{
				products.ErrProductNotFound,
				http.StatusNotFound, "product_not_found", "Product not found", []interface{}{},
			},
		},
	}
}
