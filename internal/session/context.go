package session

import "context"

type contextKey string

const userContextKey = contextKey("user")

func GetUserSession(ctx context.Context) *UserSession {
	user, _ := ctx.Value(userContextKey).(*UserSession)
	return user
}

func WithUserSession(ctx context.Context, user *UserSession) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}
