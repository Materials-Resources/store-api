package handler

import (
	"context"
	"fmt"
	"github.com/go-faster/errors"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/session"
	"github.com/ogen-go/ogen/ogenerrors"
)

var _ oas.SecurityHandler = (*SecurityHandler)(nil)

func NewSecurityHandler() *SecurityHandler {
	parser := session.NewParser("https://auth.materials-resources.com/oauth/v2/keys")
	return &SecurityHandler{
		parser: *parser,
	}
}

type SecurityHandler struct {
	parser session.Parser
}

func (s SecurityHandler) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {

	fmt.Println("handling bearer auth")
	if t.Token == "" {
		fmt.Println("no token")
		// Signal the server to skip security handling
		return ctx, nil
	}

	_, err := s.parser.ParseJwt(t.GetToken())

	// TODO custom implementation of security error
	if err != nil {
		return nil, &ogenerrors.SecurityError{
			Err: errors.New("Could not parse token"),
		}
	}

	return ctx, nil

}
