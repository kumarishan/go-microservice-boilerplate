package products

import "github.com/kumarishan/go-microservice-boilerplate/internal/products"

func MapProductsToProductDtos(products []*products.Product) []*ProductDto {
	var dtos []*ProductDto
	for _, product := range products {
		dtos = append(dtos, MapProductToProductDto(product))
	}
	return dtos
}

func MapProductToProductDto(product *products.Product) *ProductDto {
	return &ProductDto{
		Id:     product.ID.String(),
		Name:   product.Name,
		Status: product.Status.String(),
	}
}
