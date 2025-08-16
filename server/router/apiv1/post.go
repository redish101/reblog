package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerPostRoutes(api *route.RouterGroup) {
	postHandler := handler.NewPostHandler()

	posts := api.Group("/posts")
	{
		posts.POST("", postHandler.Create)
	}
}
