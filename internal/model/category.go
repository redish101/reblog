package model

type CategoryModel struct {
	BaseModel

	Name        string `gorm:"uniqueIndex" json:"name"`
	Description string `json:"description"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
