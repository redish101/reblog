package store

import (
	"git.liteyuki.org/redish101/reblog/internal/model"
	"gorm.io/gorm"
)

type TagStore struct{}

var Tag = &TagStore{}

// Create 创建标签
func (s *TagStore) Create(tag *model.TagModel) error {
	return DB.Create(tag).Error
}

// List 列出所有标签
func (s *TagStore) List(paginationParams PaginationParams) (*PaginationResponse[*model.TagModel], error) {
	var tags []*model.TagModel
	return Paginate(DB, &paginationParams, &tags)
}

// FindByName 通过 Name 获取标签
func (s *TagStore) FindByName(name string) (*model.TagModel, error) {
	var tag model.TagModel
	if err := DB.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// FindByNameWithTx 在事务中通过 Name 获取标签
func (s *TagStore) FindByNameWithTx(tx *gorm.DB, name string) (*model.TagModel, error) {
	var tag model.TagModel
	if err := tx.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// CreateWithTx 在事务中创建标签
func (s *TagStore) CreateWithTx(tx *gorm.DB, tag *model.TagModel) error {
	return tx.Create(tag).Error
}

// UpdateByName 通过 Name 更新标签
func (s *TagStore) UpdateByName(name string, tag *model.TagModel) error {
	return DB.Model(&model.TagModel{}).Where("name = ?", name).Updates(tag).Error
}

// DeleteByName 通过 Name 删除标签
func (s *TagStore) DeleteByName(name string) error {
	return DB.Where("name = ?", name).Delete(&model.TagModel{}).Error
}

// Count 统计标签数量
func (s *TagStore) Count() (int64, error) {
	var count int64
	if err := DB.Model(&model.TagModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
