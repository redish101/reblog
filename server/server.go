package server

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/redish101/reblog/docs"
	"github.com/redish101/reblog/internal/copyright"
	"github.com/redish101/reblog/internal/env"
	"github.com/redish101/reblog/internal/ipfs"
	"github.com/redish101/reblog/internal/store"
	"github.com/redish101/reblog/server/router/apiv1"
	"github.com/sirupsen/logrus"
)

//	@title						reblog api
//	@version					1.0
//	@description				reblog 后端 api
//	@basePath					/api/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							cookie
//	@name						token
//	@description				JWT token stored in cookie
//
// 启动服务端
func Start() {
	env.Init()

	if env.Dev {
		logrus.SetLevel(logrus.DebugLevel)
		hlog.SetLevel(hlog.LevelDebug)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		hlog.SetLevel(hlog.LevelInfo)
	}

	if err := store.Init(); err != nil {
		logrus.Fatalf("[STORE] 数据库连接失败: %v", err)
	}

	ipfs.Init()

	copyright.Init()

	logrus.Infoln("[HTTP] 正在启动服务")

	h := server.Default(
		server.WithHostPorts(":" + strconv.Itoa(env.Port)),
	)

	apiv1.RegisterRoutes(h)

	h.Spin()
}
