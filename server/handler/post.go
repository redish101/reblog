package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"git.liteyuki.org/redish101/reblog/internal/model"
	"git.liteyuki.org/redish101/reblog/internal/store"
	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

type PostHandler struct{}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

// findOrCreateCategory 查找或创建分类
func (h *PostHandler) findOrCreateCategory(categoryName string) (*model.CategoryModel, error) {
	if categoryName == "" {
		return nil, nil // 如果没有指定分类，则返回 0
	}

	category, err := store.Category.FindByName(categoryName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCategory := &model.CategoryModel{Name: categoryName}
		if err := store.Category.Create(newCategory); err != nil {
			return nil, err
		}
		return newCategory, nil
	} else if err != nil {
		return nil, err
	}
	return category, nil
}

// findOrCreateTags 查找或创建标签
func (h *PostHandler) findOrCreateTags(tagNames []string) ([]model.TagModel, error) {
	var tags []model.TagModel
	for _, tagName := range tagNames {
		tag, err := store.Tag.FindByName(tagName)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果标签不存在，则创建新标签
			newTag := &model.TagModel{Name: tagName}
			if err := store.Tag.Create(newTag); err != nil {
				return nil, err
			}
			tags = append(tags, *newTag)
		} else if err != nil {
			return nil, err
		} else {
			tags = append(tags, *tag)
		}
	}
	return tags, nil
}

type CreatePostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Slug     string   `json:"slug"`
	Summary  string   `json:"summary"`
	Category string   `json:"category"`
	IsDraft  bool     `json:"is_draft"`
	Tags     []string `json:"tags"`
}

//	@summary		创建文章
//	@description	创建文章
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			body	body		CreatePostRequest		true	"创建文章的请求体"
//	@success		200		{object}	model.PostModel			"创建成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		401		{object}	common.FailureResponse	"未授权"
//	@failure		403		{object}	common.FailureResponse	"权限不足"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/posts [post]
//
// 创建文章
func (h *PostHandler) Create(ctx context.Context, c *app.RequestContext) {
	var req CreatePostRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	if isSlugExits, err := store.Post.IsSlugExists(req.Slug); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	} else if isSlugExits {
		common.RespFailure(c, http.StatusConflict, "Slug 已存在")
		return
	}

	// 处理分类
	category, err := h.findOrCreateCategory(req.Category)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	// 处理标签
	tags, err := h.findOrCreateTags(req.Tags)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	// 创建文章
	post := &model.PostModel{
		Title:      req.Title,
		Content:    req.Content,
		Slug:       req.Slug,
		Summary:    req.Summary,
		CategoryID: category.ID,
		Category:   *category,
		IsDraft:    req.IsDraft,
		Tags:       tags,
	}

	if err := store.Post.Create(post); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	common.RespSuccess(c, post)
}

type ListPostsRequest struct {
	store.PaginationParams

	Categories []string `json:"category"` // 分类名称
	Tags       []string `json:"tags"`     // 标签名称
}

//	@summary		列出文章
//	@description	列出文章
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			page		query		int											false	"页码，默认为 1"
//	@param			size		query		int											false	"每页数量，默认为 10"
//	@param			categories	query		[]string									false	"分类名称过滤"
//	@param			tags		query		[]string									false	"标签名称过滤"
//	@success		200			{object}	store.PaginationResponse[model.PostModel]	"获取成功"
//	@failure		400			{object}	common.FailureResponse						"请求参数错误"
//	@failure		500			{object}	common.FailureResponse						"服务器内部错误"
//	@router			/posts [get]
//
// List 列出文章
func (h *PostHandler) List(ctx context.Context, c *app.RequestContext) {
	var req ListPostsRequest

	// 解析分页参数
	if err := c.BindQuery(&req.PaginationParams); err != nil {
		common.RespBadRequest(c, "分页参数错误: "+err.Error())
		return
	}

	// 解析分类过滤参数 (支持多个分类，用逗号分隔)
	if categoriesStr := c.Query("categories"); categoriesStr != "" {
		req.Categories = strings.Split(string(categoriesStr), ",")
	}

	// 解析标签过滤参数 (支持多个标签，用逗号分隔)
	if tagsStr := c.Query("tags"); tagsStr != "" {
		req.Tags = strings.Split(string(tagsStr), ",")
	}

	// 获取当前用户信息（用于判断是否显示草稿）
	currentUser := common.GetCurrentUser(ctx)
	currentUserEmail := ""
	if currentUser != nil {
		currentUserEmail = currentUser.Email
	}

	// 调用 store 方法获取文章列表
	posts, err := store.Post.ListWithFilters(req.PaginationParams, req.Categories, req.Tags, currentUserEmail)
	if err != nil {
		common.RespInternalServerError(c, "获取文章列表失败")
		return
	}

	common.RespSuccess(c, posts)
}

//	@summary		获取文章详情
//	@description	获取文章详情
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			slug	path		string					true	"文章 Slug"
//	@success		200		{object}	model.PostModel			"获取成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		404		{object}	common.FailureResponse	"文章未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@router			/posts/{slug} [get]
//
// Get 获取文章详情
func (h *PostHandler) Get(ctx context.Context, c *app.RequestContext) {
	slug := c.Param("slug")
	if slug == "" {
		common.RespBadRequest(c, "Slug 不能为空")
		return
	}

	post, err := store.Post.FindBySlug(slug, common.GetCurrentUser(ctx).Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "文章未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取文章失败")
		return
	}

	common.RespSuccess(c, post)
}

//	@summary		更新文章
//	@description	更新文章
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			slug	path		string					true	"文章 Slug"
//	@param			body	body		CreatePostRequest		true	"更新文章的请求体"
//	@success		200		{object}	model.PostModel			"更新成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		404		{object}	common.FailureResponse	"文章未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/posts/{slug} [put]
//
// Update 更新文章
func (h *PostHandler) Update(ctx context.Context, c *app.RequestContext) {
	slug := c.Param("slug")
	if slug == "" {
		common.RespBadRequest(c, "Slug 不能为空")
		return
	}

	var req CreatePostRequest
	if err := c.BindAndValidate(&req); err != nil {
		common.RespBadRequest(c, err.Error())
		return
	}

	if isSlugExits, err := store.Post.IsSlugExists(req.Slug); err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	} else if isSlugExits && req.Slug != slug {
		common.RespFailure(c, http.StatusConflict, "Slug 已存在")
		return
	}

	post, err := store.Post.FindBySlug(slug, common.GetCurrentUser(ctx).Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "文章未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取文章失败")
		return
	}

	// 处理分类
	category, err := h.findOrCreateCategory(req.Category)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	// 处理标签
	tags, err := h.findOrCreateTags(req.Tags)
	if err != nil {
		common.RespInternalServerError(c, err.Error())
		return
	}

	// 更新文章内容
	post.Title = req.Title
	post.Content = req.Content
	post.Slug = req.Slug
	post.Summary = req.Summary
	post.CategoryID = category.ID
	post.Category = *category
	post.IsDraft = req.IsDraft
	post.Tags = tags

	if err := store.Post.Update(post); err != nil {
		common.RespInternalServerError(c, "更新文章失败")
		return
	}

	common.RespSuccess(c, post)
}


//	@summary		删除文章
//	@description	删除文章
//	@tags			post
//	@accept			json
//	@produce		json
//	@param			slug	path	string	true	"文章 Slug"
//	@success		204		"删除成功"
//	@failure		400		{object}	common.FailureResponse	"请求参数错误"
//	@failure		404		{object}	common.FailureResponse	"文章未找到"
//	@failure		500		{object}	common.FailureResponse	"服务器内部错误"
//	@security		ApiKeyAuth
//	@router			/posts/{slug} [delete]
//
// Delete 删除文章
func (h *PostHandler) Delete(ctx context.Context, c *app.RequestContext) {
	slug := c.Param("slug")
	if slug == "" {
		common.RespBadRequest(c, "Slug 不能为空")
		return
	}

	_, err := store.Post.FindBySlug(slug, common.GetCurrentUser(ctx).Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		common.RespFailure(c, http.StatusNotFound, "文章未找到")
		return
	} else if err != nil {
		common.RespInternalServerError(c, "获取文章失败")
		return
	}

	if err := store.Post.DeleteBySlug(slug); err != nil {
		common.RespInternalServerError(c, "删除文章失败")
		return
	}

	c.Status(http.StatusNoContent)
}
