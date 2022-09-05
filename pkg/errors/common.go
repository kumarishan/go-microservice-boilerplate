package errors

import (
	"github.com/kumarishan/errors"
)

var (
	ErrInvalidInput            = errors.New("invalid input")
	ErrInternal                = errors.New("internal error")
	ErrUnexpectedInternalState = errors.New("unexpected internal state error")
	ErrNotFound                = errors.New("not found error")
)
