package products

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const productIdRegex = "^prd_[0-9a-v]{20}$"

func ValidateProductId(id string) error {
	return validation.Validate(id,
		validation.Required,
		validation.Match(regexp.MustCompile(productIdRegex)))
}

func ValidateProductName(name string, required bool) error {
	return validation.Validate(name,
		validation.Required.When(required),
		validation.Length(3, 100))
}
