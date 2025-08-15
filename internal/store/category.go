package store

import "git.liteyuki.org/redish101/reblog/internal/model"

type CategoryStore struct{}

var Category = &CategoryStore{}

// CreateCategoryParams 创建分类的参数
type CreateCategoryParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

// Create 创建分类
func (s *CategoryStore) Create(params CreateCategoryParams) error {
	return db.Create(&model.CategoryModel{
		Name:        params.Name,
		Description: params.Description,
		Slug:        params.Slug,
	}).Error
}

// List 列出所有分类
func (s *CategoryStore) List(paginationParams PaginationParams) (*PaginationResponse[*model.CategoryModel], error) {
	var categories []*model.CategoryModel
	return Paginate(db, &paginationParams, &categories)
}

// FindBySlug 通过 Slug 获取分类
func (s *CategoryStore) FindBySlug(slug string) (*model.CategoryModel, error) {
	var category model.CategoryModel
	if err := db.Where("slug = ?", slug).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateBySlug 通过 Slug 更新分类
func (s *CategoryStore) UpdateBySlug(slug string, params CreateCategoryParams) error {
	return db.Model(&model.CategoryModel{}).Where("slug = ?", slug).Updates(&model.CategoryModel{
		Name:        params.Name,
		Description: params.Description,
		Slug:        params.Slug,
	}).Error
}

// DeleteBySlug 通过 Slug 删除分类
func (s *CategoryStore) DeleteBySlug(slug string) error {
	return db.Where("slug = ?", slug).Delete(&model.CategoryModel{}).Error
}

// Count 统计分类数量
func (s *CategoryStore) Count() (int64, error) {
	var count int64
	if err := db.Model(&model.CategoryModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
