package common

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func resp(c *app.RequestContext, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func RespSuccessWithStatus(c *app.RequestContext, statusCode int, data any) {
	resp(c, statusCode, data)
}

func RespSuccess(c *app.RequestContext, data any) {
	RespSuccessWithStatus(c, http.StatusOK, data)
}

type FailureResponse struct {
	Error string `json:"error"`
}

func RespFailure(c *app.RequestContext, statusCode int, message string) {
	resp(c, statusCode, FailureResponse{Error: message})
}

func RespBadRequest(c *app.RequestContext, message string) {
	RespFailure(c, http.StatusBadRequest, message)
}

func RespInternalServerError(c *app.RequestContext, message string) {
	RespFailure(c, http.StatusInternalServerError, message)
}
