package store

import (
	"gorm.io/gorm"
)

// PaginationParams 分页请求参数
type PaginationParams struct {
	Page     int `json:"page" form:"page"`           // 页码，从1开始
	PageSize int `json:"page_size" form:"page_size"` // 每页大小
}

// PaginationResponse 分页响应
type PaginationResponse[T any] struct {
	Data       []T   `json:"data"`        // 数据列表
	Page       int   `json:"page"`        // 当前页码
	PageSize   int   `json:"page_size"`   // 每页大小
	Total      int64 `json:"total"`       // 总记录数
	TotalPages int   `json:"total_pages"` // 总页数
	HasNext    bool  `json:"has_next"`    // 是否有下一页
	HasPrev    bool  `json:"has_prev"`    // 是否有上一页
}

// Normalize 标准化分页参数
func (p *PaginationParams) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100 // 限制最大每页大小
	}
}

// GetOffset 获取偏移量
func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// Paginate 分页查询工具函数
func Paginate[T any](db *gorm.DB, req *PaginationParams, result *[]T) (*PaginationResponse[T], error) {
	req.Normalize()

	var total int64

	// 计算总数
	if err := db.Model(new(T)).Count(&total).Error; err != nil {
		return nil, err
	}

	// 查询数据
	if err := db.Offset(req.GetOffset()).Limit(req.PageSize).Find(result).Error; err != nil {
		return nil, err
	}

	// 计算分页信息
	totalPages := int((total + int64(req.PageSize) - 1) / int64(req.PageSize))

	return &PaginationResponse[T]{
		Data:       *result,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    req.Page < totalPages,
		HasPrev:    req.Page > 1,
	}, nil
}

// PaginateWithQuery 带查询条件的分页
func PaginateWithQuery[T any](db *gorm.DB, req *PaginationParams, result *[]T, query func(*gorm.DB) *gorm.DB) (*PaginationResponse[T], error) {
	req.Normalize()

	var total int64

	// 应用查询条件并计算总数
	countDB := query(db.Model(new(T)))
	if err := countDB.Count(&total).Error; err != nil {
		return nil, err
	}

	// 应用查询条件并查询数据
	queryDB := query(db)
	if err := queryDB.Offset(req.GetOffset()).Limit(req.PageSize).Find(result).Error; err != nil {
		return nil, err
	}

	// 计算分页信息
	totalPages := int((total + int64(req.PageSize) - 1) / int64(req.PageSize))

	return &PaginationResponse[T]{
		Data:       *result,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    req.Page < totalPages,
		HasPrev:    req.Page > 1,
	}, nil
}
