package apiv1

import (
	"context"
	"net/http"

	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func RegisterRoutes(server *server.Hertz) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := server.Group("/api/v1")

	registerAuthRoutes(api)
	registerCategoryRoutes(api)
	registerTagRoutes(api)
	registerPostRoutes(api)
	registerHealthzRoutes(api)

	// 仅在开发环境下添加 swagger 路由
	if env.Dev {
		server.GET("/apidoc/*any", swagger.WrapHandler(swaggerFiles.Handler))
	}

	server.Any("/*any", func(ctx context.Context, c *app.RequestContext) {
		common.RespFailure(c, http.StatusNotFound, "未知的 api")
	})
}
