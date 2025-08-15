package handler

import (
	"context"

	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

//	@summary		健康检查
//	@description	健康检查
//	@tags			healthz
//	@accept			json
//	@produce		json
//	@success		200	{boolean}	true	"安然无恙！"
//	@router			/healthz [get]
//
// 健康检查
func (h *HealthzHandler) Get(ctx context.Context, c *app.RequestContext) {
	common.RespSuccess(c, true)
}
