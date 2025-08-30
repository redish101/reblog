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

type CategoryHandler struct{}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

type CreateOrUpdateCategoryRequest struct {
	Name        string `json:"name" vd:"len($) > 0"`
	Description string `json:"description"`
}

func (h *CategoryHandler) checkCategoryExists(name string) (bool, error) {
	_, err := store.Category.FindByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

//	@summary		创建分类
//	@description	创建新的分类
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			body	body		CreateOrUpdateCategoryRequest	true	"创建分类的请求体"
//	@success		200		{object}	model.CategoryModel				"创建成功"
//	@failure		400		{object}	common.FailureResponse			"请求参数错误"
//	@failure		500		{object}	common.FailureResponse			"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/categories [post]
//
// Create 创建新的分类
func (h *CategoryHandler) Create(ctx context.Context, c *app.RequestContext) {
	var category model.CategoryModel
	if err := c.BindAndValidate(&category); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	if exists, err := h.checkCategoryExists(category.Name); err != nil {
		common.RespInternalServerError(c, "检查分类是否存在失败")
		return
	} else if exists {
		common.RespBadRequest(c, "分类已存在")
		return
	}

	if err := store.Category.Create(&category); err != nil {
		common.RespInternalServerError(c, "创建分类失败")
		return
	}

	common.RespSuccess(c, category)
}

type ListCategoriesRequest struct {
	store.PaginationParams
}

//	@summary		获取分类列表
//	@description	获取分类列表
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			page		query		int												false	"页码，默认为 1"
//	@param			page_size	query		int												false	"每页数量，默认为 10"
//	@success		200			{object}	store.PaginationResponse[model.CategoryModel]	"获取成功"
//	@failure		400			{object}	common.FailureResponse							"请求参数错误"
//	@failure		500			{object}	common.FailureResponse							"服务器内部错误"
//	@router			/categories [get]
//
// List 列出所有分类
func (h *CategoryHandler) List(ctx context.Context, c *app.RequestContext) {
	var req ListCategoriesRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	categories, err := store.Category.List(req.PaginationParams)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	common.RespSuccess(c, categories)
}

//	@summary		获取分类详情
//	@description	根据分类 名称 获取分类详情
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			name	path		string					true	"分类名称"
//	@success		200		{object}	model.CategoryModel		"获取成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		404		{object}	common.FailureResponse	"分类未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@router			/categories/{name} [get]
//
// FindByName 根据分类名称获取分类详情
func (h *CategoryHandler) FindByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	if name == "" {
		common.RespBadRequest(c, "分类名称不能为空")
		return
	}

	category, err := store.Category.FindByName(name)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "分类未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取分类失败")
		return
	}

	common.RespSuccess(c, category)
}

//	@summary		更新分类
//	@description	更新分类信息
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			name	path		string							true	"分类名称"
//	@param			body	body		CreateOrUpdateCategoryRequest	true	"更新分类的请求体"
//	@success		200		{object}	model.CategoryModel				"更新成功"
//	@failure		400		{object}	common.FailureResponse			"请求参数错误"
//	@failure		404		{object}	common.FailureResponse			"分类未找到"
//	@failure		500		{object}	common.FailureResponse			"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/categories/{name} [put]
//
// Update 更新分类信息
func (h *CategoryHandler) Update(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	if name == "" {
		common.RespBadRequest(c, "分类名称不能为空")
		return
	}

	var req CreateOrUpdateCategoryRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	if exists, err := h.checkCategoryExists(req.Name); err != nil {
		common.RespInternalServerError(c, "检查分类是否存在失败")
		return
	} else if exists {
		common.RespBadRequest(c, "分类已存在")
		return
	}

	category, err := store.Category.FindByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "分类未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取分类失败")
		return
	}

	category.Name = req.Name
	category.Description = req.Description

	if err := store.Category.Update(category); err != nil {
		common.RespInternalServerError(c, "更新分类失败")
		return
	}

	common.RespSuccess(c, category)
}

//	@summary		删除分类
//	@description	删除指定名称的分类
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			name	path	string	true	"分类名称"
//	@success		204		"删除成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		404		{object}	common.FailureResponse	"分类未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/categories/{name} [delete]
//
// Delete 删除指定名称的分类
func (h *CategoryHandler) Delete(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	if name == "" {
		common.RespBadRequest(c, "分类名称不能为空")
		return
	}

	category, err := store.Category.FindByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "分类未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取分类失败")
		return
	}

	if err := store.Category.DeleteByName(category.Name); err != nil {
		common.RespInternalServerError(c, "删除分类失败")
		return
	}

	common.RespSuccessWithStatus(c, http.StatusNoContent, nil)
}
