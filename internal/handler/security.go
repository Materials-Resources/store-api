package handler

import (
	"context"
	"customer-api/internal/oas"
	"customer-api/internal/session"
	"fmt"
	"github.com/go-faster/errors"
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

	if t.Token == "" {
		return nil, errors.New("token is empty")
	}

	claims, err := s.parser.ParseJwt(t.GetToken())

	// TODO custom implementation of security error
	if err != nil {
		return nil, &ogenerrors.SecurityError{
			Err: errors.New("Could not parse token"),
		}
	}

	fmt.Println(claims.Metadata.BranchID)

	return ctx, nil

}
