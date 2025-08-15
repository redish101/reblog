package model

type TagModel struct {
	BaseModel

	Name  string      `gorm:"uniqueIndex" json:"name"`
	Slug  string      `gorm:"uniqueIndex" json:"slug"`
	Posts []PostModel `gorm:"many2many:post_tags" json:"posts"`
}

func (TagModel) TableName() string {
	return "tags"
}
