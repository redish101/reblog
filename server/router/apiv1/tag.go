package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"git.liteyuki.org/redish101/reblog/server/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerTagRoutes(api *route.RouterGroup) {
	tagHandler := handler.NewTagHandler()

	tagsAdmin := api.Group("/tags").Use(middleware.UseAuth(true))
	{
		tagsAdmin.POST("", tagHandler.Create)
		tagsAdmin.PUT("/:name", tagHandler.Update)
		tagsAdmin.DELETE("/:name", tagHandler.Delete)
	}

	tagsPublic := api.Group("/tags").Use(middleware.UseAuth(false))
	{
		tagsPublic.GET("", tagHandler.List)
		tagsPublic.GET("/:name", tagHandler.FindByName)
	}
}
