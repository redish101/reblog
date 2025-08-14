package apiv1

import (
	"context"
	"net/http"

	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoutes(server *server.Hertz) {
	api := server.Group("/api/v1")

	registerHealthzRoutes(api)

	server.Any("/*any", func(ctx context.Context, c *app.RequestContext) {
		common.RespFailure(c, http.StatusNotFound, "未知的 api")
	})
}
