package middleware

import (
	"context"
	"net/http"

	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/internal/jwt"
	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
)

func UseAuth(requireOwner bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Cookie(common.TokenCookieKey)
		tokenString := string(token)

		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			common.RespFailure(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// 将用户信息设置到原生 Go context 中
		currentUser := &common.CurrentUser{
			Username: claims.Username,
			Email:    claims.Email,
			Nickname: claims.Nickname,
		}

		// 只使用原生 Go context
		ctx = common.SetCurrentUser(ctx, currentUser)

		if requireOwner && claims.Email != env.OwnerEmail {
			common.RespFailure(c, http.StatusForbidden, "只有数据库才能用这个的！")
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
