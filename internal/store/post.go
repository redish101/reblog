package store

import (
	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/internal/model"
	"gorm.io/gorm"
)

type PostStore struct{}

var Post = &PostStore{}

// CreatePostParams 创建文章的参数
type CreatePostParams struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CategoryID uint   `json:"category_id"`
	IsDraft    bool   `json:"is_draft"`
	Content    string `json:"content"`
}

// Create 创建文章
func (s *PostStore) Create(params CreatePostParams) error {
	return db.Create(&model.PostModel{
		Slug:       params.Slug,
		Title:      params.Title,
		Summary:    params.Summary,
		CategoryID: params.CategoryID,
		IsDraft:    params.IsDraft,
		Content:    params.Content,
	}).Error
}

// List 列出所有文章
func (s *PostStore) List(paginationParams PaginationParams, currentUserEmail string) (*PaginationResponse[*model.PostModel], error) {
	var posts []*model.PostModel

	// 如果当前用户是 owner，返回包含草稿的文章
	if currentUserEmail == env.OwnerEmail {
		return Paginate(db, &paginationParams, &posts)
	}

	// 否则只返回已发布的文章（非草稿）
	return PaginateWithQuery(db, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
		return db.Where("is_draft = ?", false)
	})
}

// FindBySlug 通过 Slug 获取文章
func (s *PostStore) FindBySlug(slug string, currentUserEmail string) (*model.PostModel, error) {
	var post model.PostModel

	// 如果当前用户是 owner，可以查看所有文章（包括草稿）
	if currentUserEmail == env.OwnerEmail {
		if err := db.Where("slug = ?", slug).First(&post).Error; err != nil {
			return nil, err
		}
		return &post, nil
	}

	// 否则只能查看已发布的文章（非草稿）
	if err := db.Where("slug = ? AND is_draft = ?", slug, false).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdateBySlug 通过 Slug 更新文章
func (s *PostStore) UpdateBySlug(slug string, params CreatePostParams) error {
	return db.Model(&model.PostModel{}).Where("slug = ?", slug).Updates(&model.PostModel{
		Slug:       params.Slug,
		Title:      params.Title,
		Summary:    params.Summary,
		CategoryID: params.CategoryID,
		IsDraft:    params.IsDraft,
		Content:    params.Content,
	}).Error
}

// DeleteBySlug 通过 Slug 删除文章
func (s *PostStore) DeleteBySlug(slug string) error {
	return db.Where("slug = ?", slug).Delete(&model.PostModel{}).Error
}

// Count 统计文章数量
func (s *PostStore) Count() (int64, error) {
	var count int64
	if err := db.Model(&model.PostModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
