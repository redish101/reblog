package store

import "git.liteyuki.org/redish101/reblog/internal/model"

type TagStore struct{}

var Tag = &TagStore{}

// CreateTagParams 创建标签的参数
type CreateTagParams struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Create 创建标签
func (s *TagStore) Create(params CreateTagParams) error {
	return db.Create(&model.TagModel{
		Name: params.Name,
		Slug: params.Slug,
	}).Error
}

// List 列出所有标签
func (s *TagStore) List(paginationParams PaginationParams) (*PaginationResponse[*model.TagModel], error) {
	var tags []*model.TagModel
	return Paginate(db, &paginationParams, &tags)
}

// FindBySlug 通过 Slug 获取标签
func (s *TagStore) FindBySlug(slug string) (*model.TagModel, error) {
	var tag model.TagModel
	if err := db.Where("slug = ?", slug).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// UpdateBySlug 通过 Slug 更新标签
func (s *TagStore) UpdateBySlug(slug string, params CreateTagParams) error {
	return db.Model(&model.TagModel{}).Where("slug = ?", slug).Updates(&model.TagModel{
		Name: params.Name,
		Slug: params.Slug,
	}).Error
}

// DeleteBySlug 通过 Slug 删除标签
func (s *TagStore) DeleteBySlug(slug string) error {
	return db.Where("slug = ?", slug).Delete(&model.TagModel{}).Error
}

// Count 统计标签数量
func (s *TagStore) Count() (int64, error) {
	var count int64
	if err := db.Model(&model.TagModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
