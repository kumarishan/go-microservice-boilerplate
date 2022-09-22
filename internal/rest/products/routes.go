package products

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/internal/rest"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"go.uber.org/dig"
)

var _ = di.Provide(NewProductRoutes, dig.Name("api/products"))

type ProductRoutes struct {
	api        *gin.RouterGroup
	controller Controller
}

type ProductRoutesParams struct {
	dig.In

	Api        *gin.RouterGroup `name:"api"`
	Controller Controller
}

func NewProductRoutes(params ProductRoutesParams) rest.Routes {
	fmt.Println("new product routes")
	return &ProductRoutes{
		params.Api,
		params.Controller,
	}
}

func (r *ProductRoutes) SetupRoutes() {
	routes := r.api.Group("/products")
	{
		routes.GET("/", rest.GinHandler(r.controller.GetProducts))
		routes.GET("/:id", rest.GinHandler(r.controller.GetProduct))
		routes.POST("/", rest.GinHandler(r.controller.AddProduct))
	}
}
