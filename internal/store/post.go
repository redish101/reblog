package store

import (
	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/internal/model"
	"gorm.io/gorm"
)

type PostStore struct{}

var Post = &PostStore{}

// Create 创建文章
func (s *PostStore) Create(post *model.PostModel) error {
	return DB.Create(post).Preload("Tags").Preload("Category").Error
}

// CreateWithTx 在事务中创建文章
func (s *PostStore) CreateWithTx(tx *gorm.DB, post *model.PostModel) error {
	return tx.Create(post).Preload("Tags").Preload("Category").Error
}

// List 列出所有文章
func (s *PostStore) List(paginationParams PaginationParams, currentUserEmail string) (*PaginationResponse[*model.PostModel], error) {
	var posts []*model.PostModel

	// 如果当前用户是 owner，返回包含草稿的文章
	if currentUserEmail == env.OwnerEmail {
		return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
			return db.Preload("Tags").Preload("Category").Order("created_at DESC")
		})
	}

	// 否则只返回已发布的文章（非草稿）
	return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
		return db.Where("is_draft = ?", false).Preload("Tags").Preload("Category").Order("created_at DESC")
	})
}

// ListWithFilters 根据分类和标签过滤列出文章
func (s *PostStore) ListWithFilters(paginationParams PaginationParams, categories []string, tags []string, currentUserEmail string) (*PaginationResponse[*model.PostModel], error) {
	var posts []*model.PostModel

	return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
		query := db.Preload("Tags").Preload("Category")

		// 如果不是 owner，只显示已发布的文章
		if currentUserEmail != env.OwnerEmail {
			query = query.Where("is_draft = ?", false)
		}

		// 按分类过滤
		if len(categories) > 0 {
			query = query.Joins("JOIN categories ON posts.category_id = categories.id").
				Where("categories.name IN ?", categories)
		}

		// 标签过滤
		if len(tags) > 0 {
			// 使用子查询来过滤包含指定标签的文章
			query = query.Where("posts.id IN (?)",
				DB.Table("posts").
					Joins("JOIN post_tags ON posts.id = post_tags.post_model_id").
					Joins("JOIN tags ON post_tags.tag_model_id = tags.id").
					Where("tags.name IN ?", tags).
					Select("posts.id").
					Group("posts.id").
					Having("COUNT(DISTINCT tags.name) = ?", len(tags)), // 确保包含所有指定标签
			)
		}

		return query.Order("posts.created_at DESC")
	})
}

// FindBySlug 通过 Slug 获取文章
func (s *PostStore) FindBySlug(slug string, currentUserEmail string) (*model.PostModel, error) {
	var post model.PostModel

	// 如果当前用户是 owner，可以查看所有文章（包括草稿）
	if currentUserEmail == env.OwnerEmail {
		if err := DB.Where("slug = ?", slug).Preload("Tags").Preload("Category").First(&post).Error; err != nil {
			return nil, err
		}
		return &post, nil
	}

	// 否则只能查看已发布的文章（非草稿）
	if err := DB.Where("slug = ? AND is_draft = ?", slug, false).
		Preload("Tags").Preload("Category").
		First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// Update 更新文章
func (s *PostStore) Update(post *model.PostModel) error {
	return DB.Updates(post).Error
}

// DeleteBySlug 通过 Slug 删除文章
func (s *PostStore) DeleteBySlug(slug string) error {
	return DB.Where("slug = ?", slug).Delete(&model.PostModel{}).Error
}

func (s *PostStore) FindByTag(tagName string, paginationParams PaginationParams, currentUserEmail string) (*PaginationResponse[*model.PostModel], error) {
	var posts []*model.PostModel

	if currentUserEmail == env.OwnerEmail {
		// 如果是 owner，返回所有文章（包括草稿）
		return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
			return db.Preload("Tags").Preload("Category").
				Joins("JOIN post_tags ON posts.id = post_tags.post_model_id").
				Joins("JOIN tags ON post_tags.tag_model_id = tags.id").
				Where("tags.name = ?", tagName).
				Order("posts.created_at DESC")
		})
	}

	// 否则只返回已发布的文章（非草稿）
	return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tags").Preload("Category").
			Joins("JOIN post_tags ON posts.id = post_tags.post_model_id").
			Joins("JOIN tags ON post_tags.tag_model_id = tags.id").
			Where("tags.name = ? AND posts.is_draft = ?", tagName, false).
			Order("posts.created_at DESC")
	})
}

// FindByCategory 通过分类名称查找文章
func (s *PostStore) FindByCategory(categoryName string, paginationParams PaginationParams, currentUserEmail string) (*PaginationResponse[*model.PostModel], error) {
	var posts []*model.PostModel

	if currentUserEmail == env.OwnerEmail {
		// 如果是 owner，返回所有文章（包括草稿）
		return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
			return db.Preload("Tags").Preload("Category").
				Joins("JOIN categories ON posts.category_id = categories.id").
				Where("categories.name = ?", categoryName).
				Order("posts.created_at DESC")
		})
	}

	// 否则只返回已发布的文章（非草稿）
	return PaginateWithQuery(DB, &paginationParams, &posts, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tags").Preload("Category").
			Joins("JOIN categories ON posts.category_id = categories.id").
			Where("categories.name = ? AND posts.is_draft = ?", categoryName, false).
			Order("posts.created_at DESC")
	})
}

// Count 统计文章数量
func (s *PostStore) Count() (int64, error) {
	var count int64
	if err := DB.Model(&model.PostModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// IsSlugExists 检查 slug 是否已存在
func (s *PostStore) IsSlugExists(slug string) (bool, error) {
	var count int64
	if err := DB.Model(&model.PostModel{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// IsSlugExistsExcludeID 检查 slug 是否已存在（排除指定 ID，用于更新时检查）
func (s *PostStore) IsSlugExistsExcludeID(slug string, excludeID uint) (bool, error) {
	var count int64
	if err := DB.Model(&model.PostModel{}).Where("slug = ? AND id != ?", slug, excludeID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
