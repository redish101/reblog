package server

import (
	"strconv"

	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/server/router/apiv1"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sirupsen/logrus"
)

//	@title			reblog api
//	@version		1.0
//	@description	reblog 后端 api
//	@basePath		/api/v1

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

	logrus.Infoln("[HTTP] 正在启动服务")

	h := server.Default(
		server.WithHostPorts(":" + strconv.Itoa(env.Port)),
	)

	apiv1.RegisterRoutes(h)

	h.Spin()
}
