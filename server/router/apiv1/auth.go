package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerAuthRoutes(api *route.RouterGroup) {
	authHandler := handler.NewAuthHandler()

	auth := api.Group("/auth")
	{
		// GitHub OAuth 登录跳转
		auth.GET("/github", authHandler.GitHubLogin)

		// GitHub OAuth 回调
		auth.GET("/github/callback", authHandler.GitHubCallback)

		// 登出
		auth.POST("/logout", authHandler.Logout)
	}
}
