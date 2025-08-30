package apiv1

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/redish101/reblog/server/handler"
)

func registerHealthzRoutes(api *route.RouterGroup) {
	healthzHandler := handler.NewHealthzHandler()

	healthz := api.Group("/healthz")
	{
		healthz.GET("", healthzHandler.Get)
		healthz.HEAD("", healthzHandler.Get)
	}
}
