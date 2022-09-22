package products

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
)

func (r AddProductRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, products.ProductNameRule...))

}
