package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redish101/reblog/internal/env"
	"github.com/redish101/reblog/internal/jwt"
	"github.com/redish101/reblog/server/common"
)

func UseAuth(requireOwner bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Cookie(common.TokenCookieKey)
		tokenString := string(token)

		// 尝试解析 token
		claims, err := jwt.ParseToken(tokenString)

		// 如果需要管理员权限，必须验证成功
		if requireOwner {
			if err != nil {
				common.RespFailure(c, http.StatusUnauthorized, "需要登录")
				c.Abort()
				return
			}

			if claims.Email != env.OwnerEmail {
				common.RespFailure(c, http.StatusForbidden, "只有数据库才能用这个的！")
				c.Abort()
				return
			}
		}

		// 如果 token 解析成功，设置用户信息到 context
		if err == nil && claims != nil {
			currentUser := &common.CurrentUser{
				Username: claims.Username,
				Email:    claims.Email,
				Nickname: claims.Nickname,
			}
			ctx = common.SetCurrentUser(ctx, currentUser)
		}
		// 如果 token 解析失败且不需要强制认证，继续执行（context 中没有用户信息）

		c.Next(ctx)
	}
}
