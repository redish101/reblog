package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"git.liteyuki.org/redish101/reblog/server/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerPostRoutes(api *route.RouterGroup) {
	postHandler := handler.NewPostHandler()

	// 需要管理员权限的路由
	postsAdmin := api.Group("/posts").Use(middleware.UseAuth(true))
	{
		postsAdmin.POST("", postHandler.Create)
	}

	// 可选认证的路由（不登录也能访问，但登录用户获得额外权限）
	postsPublic := api.Group("/posts").Use(middleware.UseAuth(false))
	{
		postsPublic.GET("", postHandler.List)
	}
}
