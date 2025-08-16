package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerHealthzRoutes(api *route.RouterGroup) {
	healthzHandler := handler.NewHealthzHandler()

	healthz := api.Group("/healthz")
	{
		healthz.GET("", healthzHandler.Get)
		healthz.HEAD("", healthzHandler.Get)
	}
}
