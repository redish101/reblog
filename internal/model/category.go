package model

type CategoryModel struct {
	BaseModel

	Name        string `gorm:"uniqueIndex" json:"name"`
	Description string `json:"description"`
	Slug        string `gorm:"uniqueIndex" json:"slug"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
