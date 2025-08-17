package common

import (
	"context"
)

type CurrentUser struct {
	Username string
	Nickname string
	Email    string
}

// 原生 Go context 相关方法
func SetCurrentUser(ctx context.Context, currentUser *CurrentUser) context.Context {
	return context.WithValue(ctx, CurrentUserContextKey, currentUser)
}

func GetCurrentUser(ctx context.Context) *CurrentUser {
	if user, ok := ctx.Value(CurrentUserContextKey).(*CurrentUser); ok {
		return user
	}
	return nil
}

func MustGetCurrentUser(ctx context.Context) *CurrentUser {
	user := GetCurrentUser(ctx)
	if user == nil {
		panic("current user not found in context")
	}
	return user
}

func ClearCurrentUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, CurrentUserContextKey, nil)
}
