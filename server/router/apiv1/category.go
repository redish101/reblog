package apiv1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/redish101/reblog/server/handler"
	"github.com/redish101/reblog/server/middleware"
)

func registerCategoryRoutes(api *route.RouterGroup) {
	categoryHandler := handler.NewCategoryHandler()

	categoriesPublic := api.Group("/categories").Use(middleware.UseAuth(false))
	{
		categoriesPublic.GET("", categoryHandler.List)
		categoriesPublic.GET("/:name", categoryHandler.FindByName)
	}

	categoriesAdmin := api.Group("/categories").Use(middleware.UseAuth(true))
	{
		categoriesAdmin.POST("", categoryHandler.Create)
		categoriesAdmin.PUT("/:name", categoryHandler.Update)
		categoriesAdmin.DELETE("/:name", categoryHandler.Delete)
	}
}
