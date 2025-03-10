package session

import (
	"context"
	"fmt"
)

type contextKey string

const userContextKey = contextKey("user")

func (m *Manager) GetUserSession(ctx context.Context) (*UserSession, error) {
	user, ok := ctx.Value(userContextKey).(*UserSession)
	if ok != true {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

func (m *Manager) WithUserSession(ctx context.Context, user *UserSession) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}
