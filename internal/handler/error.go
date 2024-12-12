package handler

import (
	"context"
	"customer-api/internal/oas"
	"errors"
	"github.com/ogen-go/ogen/ogenerrors"
)

func (h Handler) NewError(ctx context.Context, err error) *oas.ErrorStatusCode {
	var securityError *ogenerrors.SecurityError
	switch {
	case errors.As(err, &securityError):
		return &oas.ErrorStatusCode{
			StatusCode: 401,
			Response: oas.Error{
				Code:    401,
				Message: "unauthorized",
			},
		}
	}
	return &oas.ErrorStatusCode{
		StatusCode: 500,
		Response: oas.Error{
			Code:    501,
			Message: "error occurred",
		},
	}
}
