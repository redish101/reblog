package handler

import (
	"context"
	"errors"
	"net/http"

	"git.liteyuki.org/redish101/reblog/internal/model"
	"git.liteyuki.org/redish101/reblog/internal/store"
	"git.liteyuki.org/redish101/reblog/server/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
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
		logrus.Infof("Category %s not found, creating new category", categoryName)
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
//	@param			body	body		ListPostsRequest							true	"列出文章的请求体"
//	@success		200		{object}	store.PaginationResponse[model.PostModel]	"获取成功"
//	@failure		400		{object}	common.FailureResponse						"请求参数错误"
//	@failure		500		{object}	common.FailureResponse						"服务器内部错误"
//	@router			/posts [get]
//
// 列出文章
