package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redish101/reblog/internal/model"
	"github.com/redish101/reblog/internal/store"
	"github.com/redish101/reblog/server/common"
	"gorm.io/gorm"
)

type TagHandler struct{}

func NewTagHandler() *TagHandler {
	return &TagHandler{}
}

func (h *TagHandler) checkTagExists(tagName string) (bool, error) {
	_, err := store.Tag.FindByName(tagName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

type CreateOrUpdateTagRequest struct {
	Name string `json:"name" vd:"len($) > 0"`
}

//	@summary		创建标签
//	@description	创建一个新的标签
//	@tags			tag
//	@accept			json
//	@produce		json
//	@param			body	body		CreateOrUpdateTagRequest	true	"创建标签的请求体"
//	@success		200		{object}	model.TagModel				"成功创建标签"
//	@failure		400		{object}	common.FailureResponse		"请求参数错误"
//	@failure		500		{object}	common.FailureResponse		"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/tags [post]
//
// Create 创建标签
func (h *TagHandler) Create(ctx context.Context, c *app.RequestContext) {
	var req CreateOrUpdateTagRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	if exists, err := h.checkTagExists(req.Name); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	} else if exists {
		common.RespFailure(c, http.StatusBadRequest, "标签已存在")
		return
	}

	tag := &model.TagModel{
		Name: req.Name,
	}

	if err := store.Tag.Create(tag); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	common.RespSuccess(c, tag)
}

//	@summary		更新标签
//	@description	更新现有标签
//	@tags			tag
//	@accept			json
//	@produce		json
//	@param			body	body		CreateOrUpdateTagRequest	true	"更新标签的请求体"
//	@param			name	path		string						true	"标签名称"
//	@success		200		{object}	model.TagModel				"成功更新标签"
//	@failure		400		{object}	common.FailureResponse		"请求参数错误"
//	@failure		404		{object}	common.FailureResponse		"标签未找到"
//	@failure		500		{object}	common.FailureResponse		"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/tags/{name} [put]
//
// Update 更新标签
func (h *TagHandler) Update(ctx context.Context, c *app.RequestContext) {
	var req CreateOrUpdateTagRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	name := c.Param("name")
	if name == "" {
		common.RespBadRequest(c, "标签名称不能为空")
		return
	}

	tag, err := store.Tag.FindByName(name)
	if err != nil {
		common.RespFailure(c, http.StatusNotFound, "标签未找到")
		return
	}

	if req.Name != tag.Name {
		if exists, err := h.checkTagExists(req.Name); err != nil {
			common.RespInternalServerError(c, err.Error())
			return
		} else if exists {
			common.RespFailure(c, http.StatusBadRequest, "标签已存在")
			return
		}
	}

	tag.Name = req.Name

	if err := store.Tag.Update(tag); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	common.RespSuccess(c, tag)
}

//	@summary		删除标签
//	@description	删除指定名称的标签
//	@tags			tag
//	@accept			json
//	@produce		json
//	@param			name	path	string	true	"标签名称"
//	@success		204		"成功删除标签"
//	@failure		404		{object}	common.FailureResponse	"标签未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/tags/{name} [delete]
//
// Delete 删除标签
func (h *TagHandler) Delete(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")

	if err := store.Tag.DeleteByName(name); err != nil {
		common.RespFailure(c, http.StatusNotFound, "标签未找到")
		return
	}

	common.RespSuccessWithStatus(c, http.StatusNoContent, nil)
}

//	@summary		获取标签详情
//	@description	获取指定名称的标签详情
//	@tags			tag
//	@accept			json
//	@produce		json
//	@param			name	path		string					true	"标签名称"
//	@success		200		{object}	model.TagModel			"成功获取标签详情"
//	@failure		404		{object}	common.FailureResponse	"标签未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@router			/tags/{name} [get]
//
// FindByName 获取标签详情
func (h *TagHandler) FindByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")

	tag, err := store.Tag.FindByName(name)
	if err != nil {
		common.RespFailure(c, http.StatusNotFound, "标签未找到")
		return
	}

	common.RespSuccess(c, tag)
}

type ListTagsRequest struct {
	store.PaginationParams
}

//	@summary		获取标签列表
//	@description	获取标签的列表
//	@tags			tag
//	@accept			json
//	@produce		json
//	@param			page		query		int											false	"页码，默认为 1"
//	@param			page_size	query		int											false	"每页数量，默认为 10"
//	@success		200			{object}	store.PaginationResponse[model.TagModel]	"成功获取标签列表"
//	@failure		400			{object}	common.FailureResponse						"请求参数错误"
//	@failure		500			{object}	common.FailureResponse						"服务器内部错误"
//	@router			/tags [get]
//
// List 获取标签列表
func (h *TagHandler) List(ctx context.Context, c *app.RequestContext) {
	var req ListTagsRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	tags, err := store.Tag.List(req.PaginationParams)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	common.RespSuccess(c, tags)
}
