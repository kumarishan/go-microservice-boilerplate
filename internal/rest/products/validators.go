package products

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (r AddProductRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(3, 50)))

}
