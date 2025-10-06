package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type ServerlessHandler struct{}

func NewServerlessHandler() *ServerlessHandler {
	return &ServerlessHandler{}
}

//	@summary		执行云函数
//	@description	执行云函数
//	@tags			serverless
//	@success		200		{object}	json				"云函数返回的内容"
//	@security		ApiKeyAuth
//	@router			/serverless/{name}/{path} [any]
//
// 执行云函数
func (h *ServerlessHandler) Execute(ctx context.Context, c *app.RequestContext) {
	
}