package apiv1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/redish101/reblog/server/handler"
	"github.com/redish101/reblog/server/middleware"
)

func registerAuthRoutes(api *route.RouterGroup) {
	authHandler := handler.NewAuthHandler()

	auth := api.Group("/auth")
	{
		// GitHub OAuth 登录跳转
		auth.GET("/github", authHandler.GitHubLogin)

		// GitHub OAuth 回调
		auth.GET("/github/callback", authHandler.GitHubCallback)

		// 获取当前用户信息（需要认证）
		auth.GET("/me", middleware.UseAuth(false), authHandler.GetCurrentUser)

		// 登出
		auth.POST("/logout", authHandler.Logout)
	}
}
