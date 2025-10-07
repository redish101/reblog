package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type ServerlessHandler struct{}

func NewServerlessHandler() *ServerlessHandler {
	return &ServerlessHandler{}
}

func (h *ServerlessHandler) Execute(ctx context.Context, c *app.RequestContext) {
}
