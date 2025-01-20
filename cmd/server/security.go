package main

import (
	"context"
	"fmt"
	"github.com/materials-resources/customer-api/internal/oas"
	"github.com/materials-resources/customer-api/internal/session"
)

var _ oas.SecurityHandler = (*SecurityHandler)(nil)

func NewSecurityHandler(sessionManager *session.Manager) *SecurityHandler {
	return &SecurityHandler{
		sessionManager: sessionManager,
	}
}

type SecurityHandler struct {
	sessionManager *session.Manager
}

func (s SecurityHandler) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	if t.GetToken() == "" {
		ctx = s.sessionManager.WithUserSession(ctx, session.AnonymousUserSession)
		return ctx, nil
	}

	userSession := &session.UserSession{
		AccessToken: t.GetToken(),
	}

	claims, err := s.sessionManager.ParseJwt(t.GetToken())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userSession.Profile = &session.Profile{
		ContactID: claims.Metadata.ContactID,
		BranchID:  claims.Metadata.BranchID,
	}

	ctx = s.sessionManager.WithUserSession(ctx, userSession)

	return ctx, nil

}
