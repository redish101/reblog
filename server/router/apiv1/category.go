package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"git.liteyuki.org/redish101/reblog/server/middleware"
	"github.com/cloudwego/hertz/pkg/route"
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
