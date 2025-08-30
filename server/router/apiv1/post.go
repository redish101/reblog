package apiv1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/redish101/reblog/server/handler"
	"github.com/redish101/reblog/server/middleware"
)

func registerPostRoutes(api *route.RouterGroup) {
	postHandler := handler.NewPostHandler()

	postsAdmin := api.Group("/posts").Use(middleware.UseAuth(true))
	{
		postsAdmin.POST("", postHandler.Create)
		postsAdmin.PUT("/:slug", postHandler.Update)
		postsAdmin.DELETE("/:slug", postHandler.Delete)
	}

	postsPublic := api.Group("/posts").Use(middleware.UseAuth(false))
	{
		postsPublic.GET("", postHandler.List)
		postsPublic.GET("/:slug", postHandler.FindByName)
	}
}
