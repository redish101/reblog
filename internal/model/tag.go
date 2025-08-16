package model

type TagModel struct {
	BaseModel

	Name string `gorm:"uniqueIndex" json:"name"`
}

func (TagModel) TableName() string {
	return "tags"
}
