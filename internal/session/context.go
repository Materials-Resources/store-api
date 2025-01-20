package session

import "context"

type contextKey string

const userContextKey = contextKey("user")

func (m *Manager) GetUserSession(ctx context.Context) *UserSession {
	user, _ := ctx.Value(userContextKey).(*UserSession)
	return user
}

func (m *Manager) WithUserSession(ctx context.Context, user *UserSession) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}
