package apiv1

import (
	"git.liteyuki.org/redish101/reblog/server/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerHealthzRoutes(group *route.RouterGroup) {
	healthzHandler := handler.NewHealthzHandler()

	healthz := group.Group("/healthz")
	{
		healthz.GET("", healthzHandler.Get)
		healthz.HEAD("", healthzHandler.Get)
	}
}
