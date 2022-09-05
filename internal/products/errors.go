package products

import "github.com/kumarishan/errors"

var (
	ErrProductNotFound = errors.Extend(errors.ErrNotFound, "product not found")
)
