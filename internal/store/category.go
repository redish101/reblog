package store

import (
	"git.liteyuki.org/redish101/reblog/internal/model"
	"gorm.io/gorm"
)

type CategoryStore struct{}

var Category = &CategoryStore{}

// Create 创建分类
func (s *CategoryStore) Create(category *model.CategoryModel) error {
	return DB.Create(category).Error
}

// List 列出所有分类
func (s *CategoryStore) List(paginationParams PaginationParams) (*PaginationResponse[*model.CategoryModel], error) {
	var categories []*model.CategoryModel
	return Paginate(DB, &paginationParams, &categories)
}

// Update 更新分类
func (s *CategoryStore) Update(category *model.CategoryModel) error {
	return DB.Updates(category).Error
}

// FindByName 通过 Name 获取分类
func (s *CategoryStore) FindByName(name string) (*model.CategoryModel, error) {
	var category model.CategoryModel
	if err := DB.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// FindByNameWithTx 在事务中通过 Name 获取分类
func (s *CategoryStore) FindByNameWithTx(tx *gorm.DB, name string) (*model.CategoryModel, error) {
	var category model.CategoryModel
	if err := tx.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateWithTx 在事务中创建分类
func (s *CategoryStore) CreateWithTx(tx *gorm.DB, category *model.CategoryModel) error {
	return tx.Create(category).Error
}

// UpdateByName 通过 Name 更新分类
func (s *CategoryStore) UpdateByName(name string, category *model.CategoryModel) error {
	return DB.Model(&model.CategoryModel{}).Where("name = ?", name).Updates(category).Error
}

// DeleteByName 通过 Name 删除分类
func (s *CategoryStore) DeleteByName(name string) error {
	return DB.Where("name = ?", name).Delete(&model.CategoryModel{}).Error
}

// Count 统计分类数量
func (s *CategoryStore) Count() (int64, error) {
	var count int64
	if err := DB.Model(&model.CategoryModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
