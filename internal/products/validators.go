package products

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const productIdRegex = "^[0-9a-v]{20}$"

var ProductNameRule = []validation.Rule{
	validation.Required,
	validation.Length(3, 100),
}

func ValidateProductId(id string) error {
	return validation.Validate(id,
		validation.Required,
		validation.Match(regexp.MustCompile(productIdRegex)))
}

func ValidateProductName(name string) error {
	return validation.Validate(name, ProductNameRule...)
}
